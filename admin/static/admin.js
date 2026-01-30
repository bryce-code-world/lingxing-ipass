function qs(id) { return document.getElementById(id); }

function showToast(msg) {
  const el = qs("toast");
  if (!el) return;
  el.textContent = msg;
  el.style.display = "block";
  clearTimeout(window.__toastTimer);
  window.__toastTimer = setTimeout(() => { el.style.display = "none"; }, 3500);
}

async function apiJSON(method, url, body) {
  const opts = { method, headers: { "Content-Type": "application/json" } };
  if (body !== undefined) opts.body = JSON.stringify(body);
  const res = await fetch(url, opts);
  const data = await res.json().catch(() => null);
  if (!res.ok || (data && data.code !== 0)) {
    const msg = (data && data.message) ? data.message : (res.status + " " + res.statusText);
    throw new Error(msg);
  }
  return data.data;
}

async function apiDownload(url, body, filenameHint) {
  const res = await fetch(url, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify(body || {}) });
  if (!res.ok) throw new Error(res.status + " " + res.statusText);
  const blob = await res.blob();
  const cd = res.headers.get("content-disposition") || "";
  let filename = filenameHint || "export.csv";
  const m = /filename=\"?([^\";]+)\"?/i.exec(cd);
  if (m && m[1]) filename = m[1];
  const a = document.createElement("a");
  a.href = URL.createObjectURL(blob);
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  a.remove();
  setTimeout(() => URL.revokeObjectURL(a.href), 2000);
}

function adminReload() { location.reload(); }

async function adminRunAllJobs() {
  try {
    await apiJSON("POST", "/admin/api/jobs/run", {});
    showToast("Run all jobs: OK");
  } catch (e) {
    showToast("Run all jobs: " + e.message);
  }
}

async function adminRunOneJob(job) {
  try {
    await apiJSON("POST", "/admin/api/jobs/run_one", { job });
    showToast("Run " + job + ": OK");
  } catch (e) {
    showToast("Run " + job + ": " + e.message);
  }
}

async function adminRunOneJobFromInput() {
  const job = (qs("oneJob")?.value || "").trim();
  if (!job) return showToast("missing job");
  return adminRunOneJob(job);
}

async function adminManualPull() {
  try {
    const start = parseInt((qs("pullStart")?.value || "").trim(), 10);
    const end = parseInt((qs("pullEnd")?.value || "").trim(), 10);
    const status = parseInt((qs("pullStatus")?.value || "").trim(), 10);
    await apiJSON("POST", "/admin/api/dsco_order_sync/pull", { start, end, status });
    showToast("Manual pull: OK");
  } catch (e) {
    showToast("Manual pull: " + e.message);
  }
}

async function adminLoadConfig() {
  try {
    const domain = (qs("cfgDomain")?.value || "").trim() || "dsco_lingxing";
    const data = await apiJSON("GET", "/admin/api/config/runtime?domain=" + encodeURIComponent(domain));
    qs("cfgDomain").value = data.domain;
    qs("cfgJson").value = JSON.stringify(data.config, null, 2);
    showToast("Load: OK");
  } catch (e) {
    showToast("Load: " + e.message);
  }
}

async function adminSaveConfig() {
  try {
    const domain = (qs("cfgDomain")?.value || "").trim() || "dsco_lingxing";
    const raw = qs("cfgJson")?.value || "{}";
    const cfg = JSON.parse(raw);
    await apiJSON("PUT", "/admin/api/config/runtime?domain=" + encodeURIComponent(domain), cfg);
    showToast("Save: OK");
  } catch (e) {
    showToast("Save: " + e.message);
  }
}

// Orders page
let ordersOffset = 0;
let ordersLimit = 50;
let ordersTotal = 0;

function ordersFilter() {
  const f = {};
  const set = (k, v) => { if (v !== "" && v !== null && v !== undefined) f[k] = v; };
  set("po_number", (qs("qPo")?.value || "").trim());
  set("dsco_order_id", (qs("qDscoOrderId")?.value || "").trim());
  set("consumer_order_number", (qs("qConsumer")?.value || "").trim());
  set("channel", (qs("qChannel")?.value || "").trim());
  set("msku", (qs("qMSKU")?.value || "").trim());
  set("status", (qs("qStatus")?.value || "").trim());
  set("start", (qs("qStart")?.value || "").trim());
  set("end", (qs("qEnd")?.value || "").trim());
  return f;
}

async function adminLoadOrders(offset) {
  try {
    ordersOffset = offset;
    const q = new URLSearchParams({ offset: String(ordersOffset), limit: String(ordersLimit), ...ordersFilter() });
    const data = await apiJSON("GET", "/admin/api/dsco_order_sync/list?" + q.toString());
    ordersTotal = data.total || 0;
    qs("ordersMeta").textContent = "total=" + ordersTotal;
    qs("ordersPage").textContent = "offset=" + ordersOffset + ", limit=" + ordersLimit;
    const tbody = qs("ordersTable")?.querySelector("tbody");
    if (!tbody) return;
    tbody.innerHTML = "";
    for (const it of (data.items || [])) {
      const tr = document.createElement("tr");
      tr.innerHTML = `<td>${it.id}</td><td><code>${it.po_number}</code></td><td>${it.dsco_create_time}</td><td>${it.status}</td><td>${it.warehouse_id}</td><td>${it.shipment}</td><td>${it.shipped_tracking_no}</td><td>${it.dsco_invoice_id}</td>`;
      tbody.appendChild(tr);
    }
    showToast("Orders: OK");
  } catch (e) {
    showToast("Orders: " + e.message);
  }
}

function adminPrevOrders() { adminLoadOrders(Math.max(0, ordersOffset - ordersLimit)); }
function adminNextOrders() { adminLoadOrders(Math.min(Math.max(0, ordersTotal - ordersLimit), ordersOffset + ordersLimit)); }

async function adminExportOrders() {
  try {
    const f = ordersFilter();
    const body = {};
    if (f.start) body.startTime = parseInt(f.start, 10);
    if (f.end) body.endTime = parseInt(f.end, 10);
    if (f.status) body.statusIn = f.status.split(",").map(s => parseInt(s.trim(), 10)).filter(n => !Number.isNaN(n));
    if (f.po_number) body.poNumberLike = f.po_number;
    if (f.dsco_order_id) body.dscoOrderId = f.dsco_order_id;
    if (f.consumer_order_number) body.consumerOrderNumberLike = f.consumer_order_number;
    if (f.channel) body.channel = f.channel;
    if (f.msku) body.msku = f.msku;
    await apiDownload("/admin/api/export/dsco_order_sync", body, "dsco_order_sync.csv");
    showToast("Export: OK");
  } catch (e) {
    showToast("Export: " + e.message);
  }
}

// Warehouses page
let whOffset = 0;
let whLimit = 50;
let whTotal = 0;

function whFilter() {
  const f = {};
  const set = (k, v) => { if (v !== "" && v !== null && v !== undefined) f[k] = v; };
  set("dsco_warehouse_id", (qs("wDscoWh")?.value || "").trim());
  set("dsco_warehouse_sku", (qs("wDscoSku")?.value || "").trim());
  set("lingxing_warehouse_id", (qs("wLxWh")?.value || "").trim());
  set("lingxing_warehouse_sku", (qs("wLxSku")?.value || "").trim());
  set("status", (qs("wStatus")?.value || "").trim());
  set("start", (qs("wStart")?.value || "").trim());
  set("end", (qs("wEnd")?.value || "").trim());
  return f;
}

async function adminLoadWarehouses(offset) {
  try {
    whOffset = offset;
    const q = new URLSearchParams({ offset: String(whOffset), limit: String(whLimit), ...whFilter() });
    const data = await apiJSON("GET", "/admin/api/dsco_warehouse_sync/list?" + q.toString());
    whTotal = data.total || 0;
    qs("whMeta").textContent = "total=" + whTotal;
    qs("whPage").textContent = "offset=" + whOffset + ", limit=" + whLimit;
    const tbody = qs("whTable")?.querySelector("tbody");
    if (!tbody) return;
    tbody.innerHTML = "";
    for (const it of (data.items || [])) {
      const tr = document.createElement("tr");
      tr.innerHTML = `<td>${it.id}</td><td>${it.sync_time}</td><td>${it.dsco_warehouse_id}</td><td>${it.dsco_warehouse_sku}</td><td>${it.dsco_warehouse_num}</td><td>${it.lingxing_warehouse_id}</td><td>${it.lingxing_warehouse_sku}</td><td>${it.lingxing_warehouse_num}</td><td>${it.status}</td><td>${it.reason}</td>`;
      tbody.appendChild(tr);
    }
    showToast("Warehouses: OK");
  } catch (e) {
    showToast("Warehouses: " + e.message);
  }
}

function adminPrevWarehouses() { adminLoadWarehouses(Math.max(0, whOffset - whLimit)); }
function adminNextWarehouses() { adminLoadWarehouses(Math.min(Math.max(0, whTotal - whLimit), whOffset + whLimit)); }

async function adminExportWarehouses() {
  try {
    const f = whFilter();
    const body = {};
    if (f.start) body.startTime = parseInt(f.start, 10);
    if (f.end) body.endTime = parseInt(f.end, 10);
    if (f.status) body.statusIn = f.status.split(",").map(s => parseInt(s.trim(), 10)).filter(n => !Number.isNaN(n));
    if (f.dsco_warehouse_id) body.dscoWarehouseID = f.dsco_warehouse_id;
    if (f.dsco_warehouse_sku) body.dscoWarehouseSKU = f.dsco_warehouse_sku;
    if (f.lingxing_warehouse_id) body.lingXingWarehouseID = f.lingxing_warehouse_id;
    if (f.lingxing_warehouse_sku) body.lingXingWarehouseSKU = f.lingxing_warehouse_sku;
    await apiDownload("/admin/api/export/dsco_warehouse_sync", body, "dsco_warehouse_sync.csv");
    showToast("Export: OK");
  } catch (e) {
    showToast("Export: " + e.message);
  }
}

