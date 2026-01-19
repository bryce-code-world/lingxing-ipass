(function(){
  function toast(msg, ok){
    var el = document.createElement('div');
    el.style.position='fixed';
    el.style.right='18px';
    el.style.bottom='18px';
    el.style.padding='10px 12px';
    el.style.border='1px solid #d2d2d7';
    el.style.borderRadius='12px';
    el.style.background= ok ? '#e8f5ff' : '#fff0f0';
    el.style.color='#1d1d1f';
    el.style.boxShadow='0 6px 20px rgba(0,0,0,.08)';
    el.style.maxWidth='420px';
    el.textContent = msg;
    document.body.appendChild(el);
    setTimeout(function(){ el.remove(); }, 2800);
  }

  async function runJob(job){
    try{
      var resp = await fetch('/admin/run?job=' + encodeURIComponent(job), {method:'POST'});
      var data = await resp.json().catch(function(){ return {}; });
      if(!resp.ok){
        toast(job + ' 失败：' + (data.error || resp.status), false);
        return;
      }
      toast(job + ' 成功', true);
    }catch(e){
      toast(job + ' 失败：' + e, false);
    }
  }

  document.addEventListener('click', function(e){
    var target = e.target;
    if(target && target.dataset && target.dataset.runJob){
      runJob(target.dataset.runJob);
    }
  });

  var watermarkForm = document.getElementById('watermarkForm');
  if(watermarkForm){
    watermarkForm.addEventListener('submit', async function(e){
      e.preventDefault();
      var job = document.getElementById('job').value.trim();
      var wm = document.getElementById('watermark').value.trim();
      if(!job){ toast('缺少 job', false); return; }
      if(!wm){ toast('缺少 watermark', false); return; }
      try{
        JSON.parse(wm);
      }catch(err){
        toast('watermark 不是合法 JSON', false); return;
      }
      try{
        var resp = await fetch('/admin/watermark/set?job=' + encodeURIComponent(job), {method:'POST', headers:{'Content-Type':'application/json'}, body: wm});
        var data = await resp.json().catch(function(){ return {}; });
        if(!resp.ok){
          toast('提交失败：' + (data.error || resp.status), false);
          return;
        }
        toast('提交成功', true);
        window.location.href = '/admin/ui/watermarks';
      }catch(err){
        toast('提交失败：' + err, false);
      }
    });
  }
})();

