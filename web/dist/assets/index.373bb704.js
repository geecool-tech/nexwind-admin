import{J as N,K as A,_ as f,M as V,u as j,v as F,N as E,w as G,O as K,P as Q}from"./index.4ad22c22.js";/* empty css               */import{d as b,C as u,D as h,aI as t,aH as e,aM as l,aN as i,G as _,E as J,bp as S,bc as $,aK as Y,aJ as H,aE as v,aF as z,bw as U,j as X,u as I,bQ as W,bR as Z,c4 as L,aY as w,bv as tt,as as et,b7 as P,bY as g,ab as ot,bo as at,be as nt,bf as st,c as x,h as D,bj as M,bX as it,bx as ct,bu as lt,c1 as _t,c2 as rt,bn as dt,c5 as ut,by as pt,b8 as mt,c6 as ft,bd as ht,b1 as vt}from"./arco.b6dedf2e.js";/* empty css                *//* empty css                *//* empty css              *//* empty css                *//* empty css               *//* empty css               */import{a as bt}from"./message.2a41c94c.js";import{u as yt}from"./loading.f776d3e0.js";/* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css               *//* empty css              *//* empty css                *//* empty css                *//* empty css                *//* empty css               *//* empty css               *//* empty css                */import"./chart.94e2d85e.js";import"./vue.02736e06.js";const gt={class:"chat-item-footer"},$t={class:"chat-item-time"},It={class:"chat-item-actions"},St={class:"chat-item-actions-item"},wt={class:"chat-item-actions-item chat-item-actions-collect"},Ct=b({__name:"chat-item",props:{itemData:{type:Object,default(){return{}}}},setup(n){return(o,r)=>{const s=S,c=N,a=A,d=$;return u(),h("div",{class:J(["chat-item",n.itemData.isCollect?"chat-item-collected":""])},[t(d,{size:4,direction:"vertical",fill:""},{default:e(()=>[t(s,{type:"warning"},{default:e(()=>[l(i(n.itemData.username),1)]),_:1}),t(s,null,{default:e(()=>[l(i(n.itemData.content),1)]),_:1}),_("div",gt,[_("div",$t,[t(s,{type:"secondary"},{default:e(()=>[l(i(n.itemData.time),1)]),_:1})]),_("div",It,[_("div",St,[t(c)]),_("div",wt,[t(a)])])])]),_:1})],2)}}});const kt=f(Ct,[["__scopeId","data-v-29116fbe"]]),xt={class:"chat-list"},Tt=b({__name:"chat-list",props:{renderList:{type:Array,default(){return[]}}},setup(n){return(o,r)=>{const s=U;return u(),h("div",xt,[(u(!0),h(Y,null,H(n.renderList,c=>(u(),v(kt,{key:c.id,"item-data":c},null,8,["item-data"]))),128)),n.renderList.length?z("",!0):(u(),v(s,{key:0,status:"404"}))])}}});const Dt=f(Tt,[["__scopeId","data-v-9cd22eeb"]]),Mt={class:"chat-panel-content"},Ft={class:"chat-panel-footer"},zt=b({__name:"chat-panel",setup(n){const{loading:o,setLoading:r}=yt(!0),s=X([]);return(async()=>{try{const{data:a}=await bt();s.value=a}catch{}finally{r(!1)}})(),(a,d)=>{const p=W,m=Z,C=L,k=V,y=w,T=$,q=tt,R=et,B=P,O=g;return u(),v(O,{class:"general-card chat-panel",title:a.$t("monitor.title.chatPanel"),bordered:!1,"header-style":{paddingBottom:"0"},"body-style":{height:"100%",paddingTop:"16px",display:"flex",flexFlow:"column"}},{default:e(()=>[t(T,{size:8},{default:e(()=>[t(m,{style:{width:"86px"},"default-value":"all"},{default:e(()=>[t(p,{value:"all"},{default:e(()=>[l(i(a.$t("monitor.chat.options.all")),1)]),_:1})]),_:1}),t(C,{placeholder:a.$t("monitor.chat.placeholder.searchCategory")},null,8,["placeholder"]),t(y,{type:"text"},{default:e(()=>[t(k)]),_:1})]),_:1}),_("div",Mt,[t(q,{loading:I(o),style:{width:"100%"}},{default:e(()=>[t(Dt,{"render-list":s.value},null,8,["render-list"])]),_:1},8,["loading"])]),_("div",Ft,[t(T,{size:8},{default:e(()=>[t(B,null,{suffix:e(()=>[t(R)]),_:1}),t(y,{type:"primary"},{default:e(()=>[l(i(a.$t("monitor.chat.update")),1)]),_:1})]),_:1})])]),_:1},8,["title"])}}});const Lt=f(zt,[["__scopeId","data-v-5f0d9791"]]),Pt=n=>(nt("data-v-7c559bf9"),n=n(),st(),n),qt={class:"studio-wrapper"},Rt=Pt(()=>_("img",{src:"http://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/c788fc704d32cf3b1136c7d45afc2669.png~tplv-uwbnlip3yd-webp.webp",class:"studio-preview"},null,-1)),Bt={class:"studio-bar"},Ot={key:0},Nt=["src"],At=b({__name:"studio",setup(n){const o=j();return(r,s)=>{const c=ot,a=at,d=S,p=$,m=g;return u(),v(m,{class:"general-card",title:r.$t("monitor.title.studioPreview")},{extra:e(()=>[t(c)]),default:e(()=>[_("div",qt,[Rt,_("div",Bt,[I(o)?(u(),h("div",Ot,[t(p,{size:12},{default:e(()=>[t(a,{size:24},{default:e(()=>[_("img",{src:I(o).avatar},null,8,Nt)]),_:1}),t(d,null,{default:e(()=>[l(i(I(o).name)+" "+i(r.$t("monitor.studioPreview.studio")),1)]),_:1})]),_:1})])):z("",!0),t(d,{type:"secondary"},{default:e(()=>[l(" 36,000 "+i(r.$t("monitor.studioPreview.watching")),1)]),_:1})])])]),_:1},8,["title"])}}});const Vt=f(At,[["__scopeId","data-v-7c559bf9"]]),jt=b({__name:"data-statistic-list",setup(n){const{t:o}=F.exports.useI18n(),r=[{cover:"http://p1-arco.byteimg.com/tos-cn-i-uwbnlip3yd/c788fc704d32cf3b1136c7d45afc2669.png~tplv-uwbnlip3yd-webp.webp",name:"\u89C6\u9891\u76F4\u64AD",duration:"00:05:19",id:"54e23ade",status:-1}],s=a=>a===-1?`<a-tag  color="red" class='data-statistic-list-cover-tag'>
            ${o("monitor.list.tag.auditFailed")}
        </a-tag>`:"",c=x(()=>[{title:o("monitor.list.title.order"),render({rowIndex:a}){const d=`<span>${a+1}</span>`;return D(M(d))}},{title:o("monitor.list.title.cover"),render({record:a}){const d=`<div class='data-statistic-list-cover-wrapper'>
              <img src=${a.cover} />
              ${s(a.status)}
            </div>`;return D(M(d))}},{title:o("monitor.list.title.name"),dataIndex:"name"},{dataIndex:"duration",title:o("monitor.list.title.duration")},{dataIndex:"id",title:o("monitor.list.title.id")}]);return(a,d)=>{const p=it,m=S;return u(),h("div",null,[t(p,{columns:c.value,data:r,"row-key":"id","row-selection":{type:"checkbox",showCheckedAll:!0},border:!1,pagination:!1},null,8,["columns"]),t(m,{type:"secondary",class:"data-statistic-list-tip"},{default:e(()=>[l(i(a.$t("monitor.list.tip.rotations"))+" "+i(r.length)+" "+i(a.$t("monitor.list.tip.rest")),1)]),_:1})])}}});const Et={class:"data-statistic-content"},Gt={class:"data-statistic-list-wrapper"},Kt={class:"data-statistic-list-header"},Qt={class:"data-statistic-list-content"},Jt=b({__name:"data-statistic",setup(n){return(o,r)=>{const s=ct,c=lt,a=_t,d=rt,p=w,m=g;return u(),v(m,{bordered:!1,"body-style":{padding:"20px"}},{default:e(()=>[t(c,{"default-active-tab":"liveMethod"},{default:e(()=>[t(s,{key:"liveMethod",title:o.$t("monitor.tab.title.liveMethod")},null,8,["title"]),t(s,{key:"onlinePopulation",title:o.$t("monitor.tab.title.onlinePopulation")},null,8,["title"])]),_:1}),_("div",Et,[t(d,{"default-value":3,type:"button"},{default:e(()=>[t(a,{value:1},{default:e(()=>[l(i(o.$t("monitor.liveMethod.normal")),1)]),_:1}),t(a,{value:2},{default:e(()=>[l(i(o.$t("monitor.liveMethod.flowControl")),1)]),_:1}),t(a,{value:3},{default:e(()=>[l(i(o.$t("monitor.liveMethod.video")),1)]),_:1}),t(a,{value:4},{default:e(()=>[l(i(o.$t("monitor.liveMethod.web")),1)]),_:1})]),_:1}),_("div",Gt,[_("div",Kt,[t(p,{type:"text"},{default:e(()=>[l(i(o.$t("monitor.editCarousel")),1)]),_:1}),t(p,{disabled:""},{default:e(()=>[l(i(o.$t("monitor.startCarousel")),1)]),_:1})]),_("div",Qt,[t(jt)])])])]),_:1})}}});const Yt=f(Jt,[["__scopeId","data-v-514de4da"]]),Ht={key:0},Ut={key:1},Xt=b({__name:"studio-status",setup(n){const{t:o}=F.exports.useI18n(),r=x(()=>[{label:"mainstream",value:"6 Mbps"},{label:o("monitor.studioStatus.frameRate"),value:"60"},{label:"hotStandby",value:"6 Mbps"},{label:o("monitor.studioStatus.frameRate"),value:"60"},{label:"coldStandby",value:"6 Mbps"},{label:o("monitor.studioStatus.frameRate"),value:"60"}]),s=x(()=>[{label:o("monitor.studioStatus.line"),value:"\u70ED\u5907"},{label:"CDN",value:"KS"},{label:o("monitor.studioStatus.play"),value:"FLV"},{label:o("monitor.studioStatus.pictureQuality"),value:"\u539F\u753B"}]);return(c,a)=>{const d=dt,p=S,m=ut,C=pt,k=g;return u(),v(k,{class:"general-card",title:c.$t("monitor.studioStatus.title.studioStatus")},{extra:e(()=>[t(d,{color:"green"},{default:e(()=>[l(i(c.$t("monitor.studioStatus.smooth")),1)]),_:1})]),default:e(()=>[t(m,{layout:"horizontal",data:r.value,column:2},{label:e(({label:y})=>[["mainstream","hotStandby","coldStandby"].includes(y)?(u(),h("span",Ht,[t(p,{style:{"padding-right":"8px"}},{default:e(()=>[l(i(c.$t(`monitor.studioStatus.${y}`)),1)]),_:2},1024),l(" "+i(c.$t("monitor.studioStatus.bitRate")),1)])):(u(),h("span",Ut,i(y),1))]),_:1},8,["data"]),t(C,{style:{"margin-bottom":"16px"},heading:6},{default:e(()=>[l(i(c.$t("monitor.studioStatus.title.pictureInfo")),1)]),_:1}),t(m,{layout:"horizontal",data:s.value,column:2},null,8,["data"])]),_:1},8,["title"])}}});const Wt=f(Xt,[["__scopeId","data-v-5fd27840"]]),Zt={};function te(n,o){const r=E,s=w,c=G,a=K,d=Q,p=$,m=g;return u(),v(m,{class:"general-card",title:n.$t("monitor.title.quickOperation")},{default:e(()=>[t(p,{direction:"vertical",fill:"",size:10},{default:e(()=>[t(s,{long:""},{icon:e(()=>[t(r)]),default:e(()=>[l(i(n.$t("monitor.quickOperation.changeClarity"))+" ",1)]),_:1}),t(s,{long:""},{icon:e(()=>[t(c)]),default:e(()=>[l(i(n.$t("monitor.quickOperation.switchStream"))+" ",1)]),_:1}),t(s,{long:""},{icon:e(()=>[t(a)]),default:e(()=>[l(i(n.$t("monitor.quickOperation.removeClarity"))+" ",1)]),_:1}),t(s,{long:""},{icon:e(()=>[t(d)]),default:e(()=>[l(i(n.$t("monitor.quickOperation.pushFlowGasket"))+" ",1)]),_:1})]),_:1})]),_:1},8,["title"])}const ee=f(Zt,[["render",te]]),oe={};function ae(n,o){const r=P,s=mt,c=ft,a=L,d=ht,p=w,m=g;return u(),v(m,{class:"general-card",title:n.$t("monitor.title.studioInfo")},{default:e(()=>[t(d,{model:{},layout:"vertical"},{default:e(()=>[t(s,{label:n.$t("monitor.studioInfo.label.studioTitle"),required:""},{default:e(()=>[t(r,{placeholder:`admin${n.$t("monitor.studioInfo.placeholder.studioTitle")}`},null,8,["placeholder"])]),_:1},8,["label"]),t(s,{label:n.$t("monitor.studioInfo.label.onlineNotification"),required:""},{default:e(()=>[t(c)]),_:1},8,["label"]),t(s,{label:n.$t("monitor.studioInfo.label.studioCategory"),required:""},{default:e(()=>[t(a)]),_:1},8,["label"]),t(s,{label:n.$t("monitor.studioInfo.label.studioCategory"),required:""},{default:e(()=>[t(a)]),_:1},8,["label"])]),_:1}),t(p,{type:"primary"},{default:e(()=>[l(i(n.$t("monitor.studioInfo.btn.fresh")),1)]),_:1})]),_:1},8,["title"])}const ne=f(oe,[["render",ae]]),se={class:"container"},ie={class:"layout"},ce={class:"layout-left-side"},le={class:"layout-content"},_e={class:"layout-right-side"},re={name:"Monitor"},de=b({...re,setup(n){return(o,r)=>{const s=vt("Breadcrumb"),c=$;return u(),h("div",se,[t(s,{items:["menu.dashboard","menu.dashboard.monitor"]}),_("div",ie,[_("div",ce,[t(Lt)]),_("div",le,[t(c,{size:16,direction:"vertical",fill:""},{default:e(()=>[t(Vt),t(Yt)]),_:1})]),_("div",_e,[t(c,{size:16,direction:"vertical",fill:""},{default:e(()=>[t(Wt),t(ee),t(ne)]),_:1})])])])}}});const Be=f(de,[["__scopeId","data-v-12acd164"]]);export{Be as default};
