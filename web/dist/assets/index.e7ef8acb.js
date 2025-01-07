import{z as j,v as F,_ as C}from"./index.4ad22c22.js";/* empty css                *//* empty css                *//* empty css                *//* empty css               */import{d as q,c as Y,C as u,D as y,aI as t,aH as o,aK as G,aJ as H,aE as w,aN as s,cd as J,ce as K,c5 as M,bc as z,j as L,aM as m,u as B,G as S,aF as R,c3 as U,aY as O,bX as X,bv as Q,bY as A,be as W,bf as Z,b1 as D,c8 as ee,c9 as te}from"./arco.b6dedf2e.js";import{u as I}from"./loading.f776d3e0.js";/* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css                *//* empty css              *//* empty css                */import"./chart.94e2d85e.js";import"./vue.02736e06.js";function T(){return j.get("/api/profile/basic")}function ae(){return j.get("/api/operation/log")}const oe={class:"item-container"},ie={key:1},le=q({__name:"profile-item",props:{type:{type:String,default:""},renderData:{type:Object,required:!0},loading:{type:Boolean,default:!1}},setup(d){const b=d,{t:a}=F.exports.useI18n(),v=Y(()=>{var p,n,_,f,l,c,r,h,$,g,k,P,x,N;const{renderData:e}=b,i=[];return i.push({title:b.type==="pre"?a("basicProfile.title.preVideo"):a("basicProfile.title.video"),data:[{label:a("basicProfile.label.video.mode"),value:((p=e==null?void 0:e.video)==null?void 0:p.mode)||"-"},{label:a("basicProfile.label.video.acquisition.resolution"),value:((n=e==null?void 0:e.video)==null?void 0:n.acquisition.resolution)||"-"},{label:a("basicProfile.label.video.acquisition.frameRate"),value:`${((_=e==null?void 0:e.video)==null?void 0:_.acquisition.frameRate)||"-"} fps`},{label:a("basicProfile.label.video.encoding.resolution"),value:((f=e==null?void 0:e.video)==null?void 0:f.encoding.resolution)||"-"},{label:a("basicProfile.label.video.encoding.rate.min"),value:`${((l=e==null?void 0:e.video)==null?void 0:l.encoding.rate.min)||"-"} bps`},{label:a("basicProfile.label.video.encoding.rate.max"),value:`${((c=e==null?void 0:e.video)==null?void 0:c.encoding.rate.max)||"-"} bps`},{label:a("basicProfile.label.video.encoding.rate.default"),value:`${((r=e==null?void 0:e.video)==null?void 0:r.encoding.rate.default)||"-"} bps`},{label:a("basicProfile.label.video.encoding.frameRate"),value:`${((h=e==null?void 0:e.video)==null?void 0:h.encoding.frameRate)||"-"} fpx`},{label:a("basicProfile.label.video.encoding.profile"),value:(($=e==null?void 0:e.video)==null?void 0:$.encoding.profile)||"-"}]}),i.push({title:b.type==="pre"?a("basicProfile.title.preAudio"):a("basicProfile.title.audio"),data:[{label:a("basicProfile.label.audio.mode"),value:((g=e==null?void 0:e.audio)==null?void 0:g.mode)||"-"},{label:a("basicProfile.label.audio.acquisition.channels"),value:`${((k=e==null?void 0:e.audio)==null?void 0:k.acquisition.channels)||"-"} ${a("basicProfile.unit.audio.channels")}`},{label:a("basicProfile.label.audio.encoding.channels"),value:`${((P=e==null?void 0:e.audio)==null?void 0:P.encoding.channels)||"-"} ${a("basicProfile.unit.audio.channels")}`},{label:a("basicProfile.label.audio.encoding.rate"),value:`${((x=e==null?void 0:e.audio)==null?void 0:x.encoding.rate)||"-"} kbps`},{label:a("basicProfile.label.audio.encoding.profile"),value:((N=e==null?void 0:e.audio)==null?void 0:N.encoding.profile)||"-"}]}),i});return(e,i)=>{const p=J,n=K,_=M,f=z;return u(),y("div",oe,[t(f,{size:16,direction:"vertical",fill:""},{default:o(()=>[(u(!0),y(G,null,H(v.value,(l,c)=>(u(),w(_,{key:c,"label-style":{textAlign:"right",width:"200px",paddingRight:"10px",color:"rgb(var(--gray-8))"},"value-style":{width:"400px"},title:l.title,data:l.data},{value:o(({value:r})=>[d.loading?(u(),w(n,{key:0,animation:!0},{default:o(()=>[t(p,{widths:["200px"],rows:1})]),_:1})):(u(),y("span",ie,s(r),1))]),_:2},1032,["title","data"]))),128))]),_:1})])}}});const V=C(le,[["__scopeId","data-v-cc041ee9"]]),E=d=>(W("data-v-b01f27a7"),d=d(),Z(),d),ne={key:0},se=E(()=>S("span",{class:"circle"},null,-1)),ce={key:1},re=E(()=>S("span",{class:"circle pass"},null,-1)),ue=q({__name:"operation-log",setup(d){const{loading:b,setLoading:a}=I(!0),v=L([]);return(async()=>{try{const{data:i}=await ae();v.value=i}catch{}finally{a(!1)}})(),(i,p)=>{const n=U,_=O,f=X,l=Q,c=A;return u(),w(c,{class:"general-card"},{title:o(()=>[m(s(i.$t("basicProfile.title.operationLog")),1)]),default:o(()=>[t(l,{loading:B(b),style:{width:"100%"}},{default:o(()=>[t(f,{data:v.value},{columns:o(()=>[t(n,{title:i.$t("basicProfile.column.contentNumber"),"data-index":"contentNumber"},null,8,["title"]),t(n,{title:i.$t("basicProfile.column.updateContent"),"data-index":"updateContent"},null,8,["title"]),t(n,{title:i.$t("basicProfile.column.status"),"data-index":"status"},{cell:o(({record:r})=>[r.status===0?(u(),y("p",ne,[se,S("span",null,s(i.$t("basicProfile.cell.auditing")),1)])):R("",!0),r.status===1?(u(),y("p",ce,[re,S("span",null,s(i.$t("basicProfile.cell.pass")),1)])):R("",!0)]),_:1},8,["title"]),t(n,{title:i.$t("basicProfile.column.updateTime"),"data-index":"updateTime"},null,8,["title"]),t(n,{title:i.$t("basicProfile.column.operation")},{cell:o(()=>[t(_,{type:"text"},{default:o(()=>[m(s(i.$t("basicProfile.cell.view")),1)]),_:1})]),_:1},8,["title"])]),_:1},8,["data"])]),_:1},8,["loading"])]),_:1})}}});const de=C(ue,[["__scopeId","data-v-b01f27a7"]]),pe={class:"container"},_e={name:"Basic"},fe=q({..._e,setup(d){const{loading:b,setLoading:a}=I(!0),{loading:v,setLoading:e}=I(!0),i=L({}),p=L({}),n=L(1),_=async()=>{try{const{data:l}=await T();i.value=l,n.value=2}catch{}finally{a(!1)}},f=async()=>{try{const{data:l}=await T();p.value=l}catch{}finally{e(!1)}};return _(),f(),(l,c)=>{const r=D("Breadcrumb"),h=O,$=z,g=ee,k=te,P=A;return u(),y("div",pe,[t(r,{items:["menu.profile","menu.profile.basic"]}),t($,{direction:"vertical",size:16,fill:""},{default:o(()=>[t(P,{class:"general-card",title:l.$t("basicProfile.title.form")},{extra:o(()=>[t($,null,{default:o(()=>[t(h,null,{default:o(()=>[m(s(l.$t("basicProfile.cancel")),1)]),_:1}),t(h,{type:"primary"},{default:o(()=>[m(s(l.$t("basicProfile.goBack")),1)]),_:1})]),_:1})]),default:o(()=>[t(k,{current:n.value,"onUpdate:current":c[0]||(c[0]=x=>n.value=x),"line-less":"",class:"steps"},{default:o(()=>[t(g,null,{default:o(()=>[m(s(l.$t("basicProfile.steps.commit")),1)]),_:1}),t(g,null,{default:o(()=>[m(s(l.$t("basicProfile.steps.approval")),1)]),_:1}),t(g,null,{default:o(()=>[m(s(l.$t("basicProfile.steps.finish")),1)]),_:1})]),_:1},8,["current"])]),_:1},8,["title"]),t(P,{class:"general-card"},{default:o(()=>[t(V,{loading:B(b),"render-data":i.value},null,8,["loading","render-data"])]),_:1}),t(P,{class:"general-card"},{default:o(()=>[t(V,{loading:B(v),type:"pre","render-data":p.value},null,8,["loading","render-data"])]),_:1}),t(de)]),_:1})])}}});const Re=C(fe,[["__scopeId","data-v-e83af96b"]]);export{Re as default};
