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
    if (e.key === "Escape") {
      hideToast();
      closeModal("editStatusModal");
    }
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

function tzForIntl(tz) {
  // 兼容：部分运行环境不接受 "UTC"，但接受 "Etc/UTC"。
  const s = (tz || "").trim();
  if (s === "UTC") return "Etc/UTC";
  return s || "Etc/UTC";
}

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
      timeZone: tzForIntl(ADMIN_TZ),
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
      timeZone: "Etc/UTC",
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
      timeZone: tzForIntl(tz),
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
      timeZone: "Etc/UTC",
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
  normalizeConfigTextarea();
});

const JOBS_DISPLAY_ORDER = [
  "pull_dsco_orders",
  "push_to_lingxing",
  "ack_to_dsco",
  "ship_to_dsco",
  "invoice_to_dsco",
  "sync_stock",
  "pull_sku_pair",
  "cleanup_exports",
];

const JOBS_WITH_MULTI_BAN = new Set(["ack_to_dsco", "ship_to_dsco", "invoice_to_dsco"]);
const JOBS_WITH_SYNC = new Set(["sync_stock"]);

function normalizeJobConfigForDisplay(name, jc) {
  if (!jc || typeof jc !== "object") return jc;
  const out = {};
  out.enable = Object.prototype.hasOwnProperty.call(jc, "enable") ? jc.enable : false;
  if (Object.prototype.hasOwnProperty.call(jc, "cron")) out.cron = jc.cron;
  if (Object.prototype.hasOwnProperty.call(jc, "size")) out.size = jc.size;
  if (JOBS_WITH_MULTI_BAN.has(name)) {
    out.multi_ban = Object.prototype.hasOwnProperty.call(jc, "multi_ban") ? jc.multi_ban : false;
  }
  if (JOBS_WITH_SYNC.has(name)) {
    out.sync = Object.prototype.hasOwnProperty.call(jc, "sync") ? jc.sync : false;
  }
  for (const k of Object.keys(jc)) {
    if (!Object.prototype.hasOwnProperty.call(out, k)) out[k] = jc[k];
  }
  return out;
}

function normalizeJobsForDisplay(jobs) {
  if (!jobs || typeof jobs !== "object") return jobs;
  const out = {};
  for (const name of JOBS_DISPLAY_ORDER) {
    if (Object.prototype.hasOwnProperty.call(jobs, name)) {
      out[name] = normalizeJobConfigForDisplay(name, jobs[name]);
    }
  }
  const rest = Object.keys(jobs)
    .filter((k) => !JOBS_DISPLAY_ORDER.includes(k))
    .sort();
  for (const k of rest) out[k] = normalizeJobConfigForDisplay(k, jobs[k]);
  return out;
}

function normalizeRuntimeConfigForDisplay(cfg) {
  if (!cfg || typeof cfg !== "object") return cfg;
  const out = {};
  if (Object.prototype.hasOwnProperty.call(cfg, "domain")) out.domain = cfg.domain;
  if (Object.prototype.hasOwnProperty.call(cfg, "jobs")) out.jobs = normalizeJobsForDisplay(cfg.jobs);
  if (Object.prototype.hasOwnProperty.call(cfg, "mapping")) out.mapping = cfg.mapping;
  for (const k of Object.keys(cfg)) {
    if (!Object.prototype.hasOwnProperty.call(out, k)) out[k] = cfg[k];
  }
  return out;
}

function normalizeConfigTextarea() {
  const ta = qs("cfgJson");
  if (!ta) return;
  const raw = (ta.value || "").trim();
  if (!raw) return;
  try {
    const cfg = JSON.parse(raw);
    const normalized = normalizeRuntimeConfigForDisplay(cfg);
    ta.value = JSON.stringify(normalized, null, 2);
  } catch (_) {
    // ignore invalid json in editor
  }
}

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
    await apiJSON("POST", "/admin/api/dsco_order_sync/pull", { start, end });
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
    qs("cfgJson").value = JSON.stringify(normalizeRuntimeConfigForDisplay(data.config), null, 2);
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
  set("warehouse_id", (qs("qWarehouse")?.value || "").trim());
  set("shipment", (qs("qShipment")?.value || "").trim());
  set("tracking", (qs("qTracking")?.value || "").trim());
  set("invoice_id", (qs("qInvoice")?.value || "").trim());
  set("dsco_status", (qs("qDscoStatus")?.value || "").trim());
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
      const poJSON = JSON.stringify(it.po_number || "");
      const skus = Array.isArray(it.mskus) ? it.mskus.filter(s => String(s || "").trim() !== "") : [];
      const skuText = skus.join(", ");
      const skuShow = skuText.length > 40 ? (skuText.slice(0, 40) + "...") : skuText;
      const tr = document.createElement("tr");
      const canRun = Number(it.status) >= 1 && Number(it.status) <= 4;
      const runBtn = canRun
        ? `<button class="btn" onclick='adminRunOneOrderByStatus(${poJSON}, ${it.status})'>Run</button>`
        : `<button class="btn" disabled title="status=5/6 disabled">Run</button>`;
      tr.innerHTML = `<td>${it.id}</td><td><code>${it.po_number}</code></td><td title="${it.dsco_create_time}">${fmtUnixSec(it.dsco_create_time)}</td><td title="${it.created_at}">${fmtUnixSec(it.created_at)}</td><td title="${it.updated_at}">${fmtUnixSec(it.updated_at)}</td><td>${it.dsco_status || ""}</td><td>${it.status}</td><td>${it.warehouse_id}</td><td>${it.shipment}</td><td>${it.dsco_retailer_id || ""}</td><td title="${skuText.replace(/\"/g, '&quot;')}"><code>${skuShow}</code></td><td>${it.shipped_tracking_no}</td><td>${it.dsco_invoice_id}</td><td><div class="actions"><button class="btn" onclick="adminViewOrderDetail(${it.id})">View</button>${runBtn}<button class="btn" onclick='adminOpenEditOrderStatus(${poJSON}, ${it.status})'>Edit</button></div></td>`;
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
    if (f.dsco_status) body.dscoStatus = f.dsco_status;
    if (f.po_number) body.poNumberLike = f.po_number;
    if (f.dsco_retailer_id) body.dscoRetailerId = f.dsco_retailer_id;
    if (f.msku) body.msku = f.msku;
    if (f.warehouse_id) body.warehouseId = f.warehouse_id;
    if (f.shipment) body.shipment = f.shipment;
    if (f.tracking) body.tracking = f.tracking;
    if (f.invoice_id) body.invoiceId = f.invoice_id;
    await apiDownload("/admin/api/export/dsco_order_sync", body, "dsco_order_sync.csv");
    showToast("Export: OK");
  } catch (e) {
    showToast("Export: " + e.message);
  }
}

function openModal(id) {
  const m = qs(id);
  if (!m) return;
  m.style.display = "block";
}

function closeModal(id) {
  const m = qs(id);
  if (!m) return;
  m.style.display = "none";
}

async function copyToClipboard(text) {
  const s = String(text || "");
  if (!s) return;
  if (navigator.clipboard && navigator.clipboard.writeText) {
    await navigator.clipboard.writeText(s);
    return;
  }
  const ta = document.createElement("textarea");
  ta.value = s;
  ta.style.position = "fixed";
  ta.style.left = "-9999px";
  document.body.appendChild(ta);
  ta.select();
  document.execCommand("copy");
  ta.remove();
}

function statusLabel(status) {
  switch (Number(status)) {
    case 1: return "1 待同步（推单到领星）";
    case 2: return "2 待确认（回传 ACK）";
    case 3: return "3 待发货回传（已确认）";
    case 4: return "4 待发票回传（已发货）";
    case 5: return "5 完成（已回传发票）";
    case 6: return "6 已取消";
    default: return String(status);
  }
}

function jobForOrderStatus(status) {
  switch (Number(status)) {
    case 1: return "push_to_lingxing";
    case 2: return "ack_to_dsco";
    case 3: return "ship_to_dsco";
    case 4: return "invoice_to_dsco";
    default: return "";
  }
}

async function adminRunOneOrderByStatus(poNumber, currentStatus) {
  const po = String(poNumber || "").trim();
  if (!po) return showToast("missing po_number");

  const st = Number(currentStatus);
  const job = jobForOrderStatus(st);
  if (!job) return showToast("Run: unsupported status=" + st);
  if (st === 5 || st === 6) return showToast("Run: disabled for status=" + st);

  if (!confirm(`Run ${job} for ${po} ?`)) return;

  try {
    const data = await apiJSON("POST", "/admin/api/dsco_order_sync/run_one", { po_number: po });
    const b = Number(data.status_before);
    const a = Number(data.status_after);
    const changed = !!data.status_changed;
    if (changed) {
      showToast(`Run: OK (${data.job}) ${b} -> ${a}`, "ok");
    } else {
      showToast(`Run: OK (${data.job}) status unchanged (${b})`, "ok");
    }
    await adminLoadOrders(ordersOffset);
  } catch (e) {
    showToast("Run: " + e.message);
  }
}

function adminOpenEditOrderStatus(poNumber, currentStatus) {
  const po = String(poNumber || "").trim();
  if (!po) return showToast("missing po_number");

  const poEl = qs("editStatusPo");
  const sel = qs("editStatusSelect");
  const tip = qs("editStatusTip");
  if (poEl) poEl.textContent = po;
  if (sel) sel.value = String(currentStatus || 1);
  if (tip) tip.textContent = "当前状态：" + statusLabel(currentStatus);
  openModal("editStatusModal");
}

async function adminViewOrderDetail(id) {
  try {
    const n = Number(id);
    if (!Number.isFinite(n) || n <= 0) return showToast("invalid id");
    const data = await apiJSON("GET", "/admin/api/dsco_order_sync/detail?id=" + encodeURIComponent(String(n)));
    qs("orderDetailTitle").textContent = "Order Detail";
    qs("orderDetailMeta").textContent = `id=${data.id}, po_number=${data.po_number || ""}`;
    const txt = JSON.stringify(data, null, 2);
    qs("orderDetailText").textContent = txt;
    openModal("orderDetailModal");
  } catch (e) {
    showToast("Detail: " + e.message);
  }
}

async function adminCopyOrderDetail() {
  try {
    const txt = qs("orderDetailText")?.textContent || "";
    await copyToClipboard(txt);
    showToast("Copied: OK");
  } catch (e) {
    showToast("Copy: " + e.message);
  }
}

async function adminSaveOrderStatus() {
  try {
    const po = (qs("editStatusPo")?.textContent || "").trim();
    const status = parseInt((qs("editStatusSelect")?.value || "").trim(), 10);
    if (!po) return showToast("missing po_number");
    if (!Number.isFinite(status) || status < 1 || status > 6) return showToast("invalid status");
    await apiJSON("PUT", "/admin/api/dsco_order_sync/status", { po_number: po, status });
    closeModal("editStatusModal");
    showToast("Update status: OK");
    await adminLoadOrders(ordersOffset);
  } catch (e) {
    showToast("Update status: " + e.message);
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
