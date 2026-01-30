function qs(id) { return document.getElementById(id); }

let __busyCount = 0;

function setButtonsDisabled(disabled) {
  for (const btn of document.querySelectorAll("button.btn")) {
    btn.disabled = !!disabled;
  }
}

function showLoading(text) {
  const m = qs("loadingModal");
  if (!m) return;
  const t = m.querySelector(".loading-text");
  if (t && text) t.textContent = text;
  m.style.display = "block";
}

function hideLoading() {
  const m = qs("loadingModal");
  if (!m) return;
  m.style.display = "none";
}

function beginBusy(label) {
  __busyCount += 1;
  if (__busyCount === 1) {
    setButtonsDisabled(true);
    showLoading(label || "Loading...");
  }
}

function endBusy() {
  __busyCount = Math.max(0, __busyCount - 1);
  if (__busyCount === 0) {
    setButtonsDisabled(false);
    hideLoading();
  }
}

function inferToastKind(msg) {
  const s = String(msg || "");
  if (s.includes(": OK") || s.trim() === "OK" || s.endsWith(" OK")) return "ok";
  return "error";
}

function showToast(msg, kind) {
  const modal = qs("toastModal");
  const text = qs("toastText");
  if (!modal || !text) return;

  const k = kind || inferToastKind(msg);
  modal.classList.remove("ok", "error");
  modal.classList.add(k === "ok" ? "ok" : "error");

  text.textContent = String(msg || "");
  modal.style.display = "block";

  clearTimeout(window.__toastTimer);
  if (k === "ok") {
    window.__toastTimer = setTimeout(() => { modal.style.display = "none"; }, 1800);
  }
}

function hideToast() {
  const modal = qs("toastModal");
  if (!modal) return;
  modal.style.display = "none";
}

function initGlobalUI() {
  const btn = qs("toastClose");
  if (btn) btn.addEventListener("click", hideToast);
  const modal = qs("toastModal");
  if (modal) {
    const mask = modal.querySelector(".modal-mask");
    if (mask) mask.addEventListener("click", hideToast);
  }
  document.addEventListener("keydown", (e) => {
    if (e.key === "Escape") hideToast();
  });
}

function normalizeTZ(raw) {
  let s = (raw || "").trim();
  if (!s) return "UTC";
  if ((s.startsWith('"') && s.endsWith('"')) || (s.startsWith("'") && s.endsWith("'"))) {
    s = s.slice(1, -1).trim();
  }
  return s || "UTC";
}

const ADMIN_TZ_RAW = (typeof window !== "undefined" && window.ADMIN_TZ) ? window.ADMIN_TZ : "UTC";
let ADMIN_TZ = normalizeTZ(ADMIN_TZ_RAW);

function setTZLabel() {
  const el = qs("tzLabel");
  if (el) el.textContent = ADMIN_TZ || "UTC";
}

function fmtUnixSec(sec) {
  const n = Number(sec);
  if (!Number.isFinite(n) || n <= 0) return "";
  const d = new Date(n * 1000);
  let dtf;
  try {
    dtf = new Intl.DateTimeFormat("sv-SE", {
      timeZone: ADMIN_TZ,
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hourCycle: "h23",
    });
  } catch (e) {
    ADMIN_TZ = "UTC";
    dtf = new Intl.DateTimeFormat("sv-SE", {
      timeZone: "UTC",
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hourCycle: "h23",
    });
  }
  return dtf.format(d).replace(" ", "T");
}

function getZonedPartsFromUTC(ms, tz) {
  const d = new Date(ms);
  let dtf;
  try {
    dtf = new Intl.DateTimeFormat("en-CA", {
      timeZone: tz,
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hourCycle: "h23",
    });
  } catch (e) {
    dtf = new Intl.DateTimeFormat("en-CA", {
      timeZone: "UTC",
      year: "numeric",
      month: "2-digit",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hourCycle: "h23",
    });
  }
  const parts = dtf.formatToParts(d);
  const out = {};
  for (const p of parts) {
    if (p.type === "year") out.year = parseInt(p.value, 10);
    if (p.type === "month") out.month = parseInt(p.value, 10);
    if (p.type === "day") out.day = parseInt(p.value, 10);
    if (p.type === "hour") out.hour = parseInt(p.value, 10);
    if (p.type === "minute") out.minute = parseInt(p.value, 10);
    if (p.type === "second") out.second = parseInt(p.value, 10);
  }
  return out;
}

function parseDateTimeLocal(v) {
  const s = (v || "").trim();
  if (!s) return null;
  const m = /^(\d{4})-(\d{2})-(\d{2})T(\d{2}):(\d{2})(?::(\d{2}))?$/.exec(s);
  if (!m) return null;
  return {
    year: parseInt(m[1], 10),
    month: parseInt(m[2], 10),
    day: parseInt(m[3], 10),
    hour: parseInt(m[4], 10),
    minute: parseInt(m[5], 10),
    second: m[6] ? parseInt(m[6], 10) : 0,
  };
}

function zonedDateTimeToUnixSec(parts, tz) {
  // Iteratively converge: find the UTC instant whose wall-clock in tz equals parts.
  let guess = Date.UTC(parts.year, parts.month - 1, parts.day, parts.hour, parts.minute, parts.second);
  for (let i = 0; i < 3; i++) {
    const got = getZonedPartsFromUTC(guess, tz);
    const desiredUTC = Date.UTC(parts.year, parts.month - 1, parts.day, parts.hour, parts.minute, parts.second);
    const gotUTC = Date.UTC(got.year, got.month - 1, got.day, got.hour, got.minute, got.second);
    const diff = desiredUTC - gotUTC;
    if (diff === 0) break;
    guess += diff;
  }
  return Math.floor(guess / 1000);
}

async function apiJSON(method, url, body) {
  beginBusy(url);
  try {
    const opts = { method, headers: { "Content-Type": "application/json" } };
    if (body !== undefined) opts.body = JSON.stringify(body);
    const res = await fetch(url, opts);
    const data = await res.json().catch(() => null);
    if (!res.ok || (data && data.code !== 0)) {
      const msg = (data && data.message) ? data.message : (res.status + " " + res.statusText);
      throw new Error(msg);
    }
    return data.data;
  } finally {
    endBusy();
  }
}

async function apiDownload(url, body, filenameHint) {
  beginBusy(url);
  try {
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
  } finally {
    endBusy();
  }
}

function adminReload() { location.reload(); }

function renderTimestamps() {
  for (const el of document.querySelectorAll(".ts[data-ts]")) {
    const raw = el.getAttribute("data-ts");
    el.textContent = fmtUnixSec(raw);
    el.title = raw || "";
  }
}

document.addEventListener("DOMContentLoaded", () => {
  initGlobalUI();
});

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
    const startP = parseDateTimeLocal(qs("pullStartDT")?.value || "");
    const endP = parseDateTimeLocal(qs("pullEndDT")?.value || "");
    const start = startP ? zonedDateTimeToUnixSec(startP, ADMIN_TZ) : 0;
    const end = endP ? zonedDateTimeToUnixSec(endP, ADMIN_TZ) : 0;
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
  set("dsco_retailer_id", (qs("qRetailer")?.value || "").trim());
  set("msku", (qs("qMSKU")?.value || "").trim());
  set("status", (qs("qStatus")?.value || "").trim());
  const s = parseDateTimeLocal(qs("qStartDT")?.value || "");
  const e = parseDateTimeLocal(qs("qEndDT")?.value || "");
  if (s) set("start", String(zonedDateTimeToUnixSec(s, ADMIN_TZ)));
  if (e) set("end", String(zonedDateTimeToUnixSec(e, ADMIN_TZ)));
  return f;
}

async function adminLoadOrders(offset) {
  try {
    const lim = parseInt((qs("qLimit")?.value || "").trim(), 10);
    if (Number.isFinite(lim) && lim > 0) ordersLimit = lim;
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
      tr.innerHTML = `<td>${it.id}</td><td><code>${it.po_number}</code></td><td title="${it.dsco_create_time}">${fmtUnixSec(it.dsco_create_time)}</td><td>${it.dsco_status || ""}</td><td>${it.status}</td><td>${it.warehouse_id}</td><td>${it.shipment}</td><td>${it.dsco_retailer_id || ""}</td><td>${it.shipped_tracking_no}</td><td>${it.dsco_invoice_id}</td>`;
      tbody.appendChild(tr);
    }
  } catch (e) {
    showToast("Orders: " + e.message);
  }
}

function adminPrevOrders() { adminLoadOrders(Math.max(0, ordersOffset - ordersLimit)); }
function adminNextOrders() { adminLoadOrders(Math.min(Math.max(0, ordersTotal - ordersLimit), ordersOffset + ordersLimit)); }

function adminSetOrdersLimit() { adminLoadOrders(0); }

async function adminExportOrders() {
  try {
    const f = ordersFilter();
    const body = {};
    if (f.start) body.startTime = parseInt(f.start, 10);
    if (f.end) body.endTime = parseInt(f.end, 10);
    if (f.status) body.statusIn = f.status.split(",").map(s => parseInt(s.trim(), 10)).filter(n => !Number.isNaN(n));
    if (f.po_number) body.poNumberLike = f.po_number;
    if (f.dsco_retailer_id) body.dscoRetailerId = f.dsco_retailer_id;
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
  const s = parseDateTimeLocal(qs("wStartDT")?.value || "");
  const e = parseDateTimeLocal(qs("wEndDT")?.value || "");
  if (s) set("start", String(zonedDateTimeToUnixSec(s, ADMIN_TZ)));
  if (e) set("end", String(zonedDateTimeToUnixSec(e, ADMIN_TZ)));
  return f;
}

async function adminLoadWarehouses(offset) {
  try {
    const lim = parseInt((qs("wLimit")?.value || "").trim(), 10);
    if (Number.isFinite(lim) && lim > 0) whLimit = lim;
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
      tr.innerHTML = `<td>${it.id}</td><td title="${it.sync_time}">${fmtUnixSec(it.sync_time)}</td><td>${it.dsco_warehouse_id}</td><td>${it.dsco_warehouse_sku}</td><td>${it.dsco_warehouse_num}</td><td>${it.lingxing_warehouse_id}</td><td>${it.lingxing_warehouse_sku}</td><td>${it.lingxing_warehouse_num}</td><td>${it.status}</td><td>${it.reason}</td>`;
      tbody.appendChild(tr);
    }
  } catch (e) {
    showToast("Warehouses: " + e.message);
  }
}

function adminPrevWarehouses() { adminLoadWarehouses(Math.max(0, whOffset - whLimit)); }
function adminNextWarehouses() { adminLoadWarehouses(Math.min(Math.max(0, whTotal - whLimit), whOffset + whLimit)); }

function adminSetWarehousesLimit() { adminLoadWarehouses(0); }

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

document.addEventListener("DOMContentLoaded", () => {
  setTZLabel();
  renderTimestamps();
  if (qs("ordersTable")) adminLoadOrders(0);
  if (qs("whTable")) adminLoadWarehouses(0);
});
