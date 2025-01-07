import{z as A,_ as V}from"./index.4ad22c22.js";/* empty css                *//* empty css                *//* empty css                */import{d as w,j as y,C as f,aE as g,aH as a,aI as t,aM as n,G as C,aN as i,b7 as T,b8 as N,bQ as q,bR as L,c7 as E,aY as S,bd as I,aW as P,c6 as D,bc as M,D as U,bw as z,by as Y,bb as j,bq as G,b1 as H,bL as K,aF as O,u as Q,c8 as W,c9 as Z,bY as J,bv as X}from"./arco.b6dedf2e.js";import{u as x}from"./loading.f776d3e0.js";/* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css               *//* empty css              *//* empty css                *//* empty css               *//* empty css               *//* empty css                *//* empty css               *//* empty css                *//* empty css                *//* empty css                */import"./chart.94e2d85e.js";import"./vue.02736e06.js";function ee(h){return A.post("/api/channel-form/submit",{data:h})}const te=w({__name:"base-info",emits:["changeStep"],setup(h,{emit:v}){const d=v,s=y(),o=y({activityName:"",channelType:"",promotionTime:[],promoteLink:"https://arco.design"}),b=async()=>{var e;await((e=s.value)==null?void 0:e.validate())||d("changeStep","forward",{...o.value})};return(r,e)=>{const l=T,m=N,p=q,_=L,$=E,F=S,k=I;return f(),g(k,{ref_key:"formRef",ref:s,model:o.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[t(m,{field:"activityName",label:r.$t("stepForm.form.label.activityName"),rules:[{required:!0,message:r.$t("stepForm.form.error.activityName.required")},{match:/^[a-zA-Z0-9\u4e00-\u9fa5]{1,20}$/,message:r.$t("stepForm.form.error.activityName.pattern")}]},{default:a(()=>[t(l,{modelValue:o.value.activityName,"onUpdate:modelValue":e[0]||(e[0]=u=>o.value.activityName=u),placeholder:r.$t("stepForm.placeholder.activityName")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(m,{field:"channelType",label:r.$t("stepForm.form.label.channelType"),rules:[{required:!0,message:r.$t("stepForm.form.error.channelType.required")}]},{default:a(()=>[t(_,{modelValue:o.value.channelType,"onUpdate:modelValue":e[1]||(e[1]=u=>o.value.channelType=u),placeholder:r.$t("stepForm.placeholder.channelType")},{default:a(()=>[t(p,null,{default:a(()=>[n("APP\u901A\u7528\u6E20\u9053")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(m,{field:"promotionTime",label:r.$t("stepForm.form.label.promotionTime"),rules:[{required:!0,message:r.$t("stepForm.form.error.promotionTime.required")}]},{default:a(()=>[t($,{modelValue:o.value.promotionTime,"onUpdate:modelValue":e[2]||(e[2]=u=>o.value.promotionTime=u)},null,8,["modelValue"])]),_:1},8,["label","rules"]),t(m,{field:"promoteLink",label:r.$t("stepForm.form.label.promoteLink"),rules:[{required:!0,message:r.$t("stepForm.form.error.promoteLink.required")},{type:"url",message:r.$t("stepForm.form.error.promoteLink.pattern")}],"row-class":"keep-margin"},{help:a(()=>[C("span",null,i(r.$t("stepForm.form.tip.promoteLink")),1)]),default:a(()=>[t(l,{modelValue:o.value.promoteLink,"onUpdate:modelValue":e[3]||(e[3]=u=>o.value.promoteLink=u),placeholder:r.$t("stepForm.placeholder.promoteLink")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(m,null,{default:a(()=>[t(F,{type:"primary",onClick:b},{default:a(()=>[n(i(r.$t("stepForm.button.next")),1)]),_:1})]),_:1})]),_:1},8,["model"])}}});const oe=V(te,[["__scopeId","data-v-b8f630c6"]]),ae=w({__name:"channel-info",emits:["changeStep"],setup(h,{emit:v}){const d=v,s=y(),o=y({advertisingSource:"",advertisingMedia:"",keyword:[],pushNotify:!0,advertisingContent:""}),b=async()=>{var l;await((l=s.value)==null?void 0:l.validate())||d("changeStep","submit",{...o.value})},r=()=>{d("changeStep","backward")};return(e,l)=>{const m=T,p=N,_=q,$=L,F=P,k=D,u=S,B=M,R=I;return f(),g(R,{ref_key:"formRef",ref:s,model:o.value,class:"form","label-col-props":{span:6},"wrapper-col-props":{span:18}},{default:a(()=>[t(p,{field:"advertisingSource",label:e.$t("stepForm.form.label.advertisingSource"),rules:[{required:!0,message:e.$t("stepForm.form.error.advertisingSource.required")}]},{default:a(()=>[t(m,{modelValue:o.value.advertisingSource,"onUpdate:modelValue":l[0]||(l[0]=c=>o.value.advertisingSource=c),placeholder:e.$t("stepForm.placeholder.advertisingSource")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(p,{field:"advertisingMedia",label:e.$t("stepForm.form.label.advertisingMedia"),rules:[{required:!0,message:e.$t("stepForm.form.error.advertisingMedia.required")}]},{default:a(()=>[t(m,{modelValue:o.value.advertisingMedia,"onUpdate:modelValue":l[1]||(l[1]=c=>o.value.advertisingMedia=c),placeholder:e.$t("stepForm.placeholder.advertisingMedia")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(p,{field:"keyword",label:e.$t("stepForm.form.label.keyword"),rules:[{required:!0,message:e.$t("stepForm.form.error.keyword.required")}]},{default:a(()=>[t($,{modelValue:o.value.keyword,"onUpdate:modelValue":l[2]||(l[2]=c=>o.value.keyword=c),placeholder:e.$t("stepForm.placeholder.keyword"),multiple:""},{default:a(()=>[t(_,null,{default:a(()=>[n("\u4ECA\u65E5\u5934\u6761")]),_:1}),t(_,null,{default:a(()=>[n("\u706B\u5C71")]),_:1})]),_:1},8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(p,{field:"pushNotify",label:e.$t("stepForm.form.label.pushNotify"),rules:[{required:!0}]},{default:a(()=>[t(F,{modelValue:o.value.pushNotify,"onUpdate:modelValue":l[3]||(l[3]=c=>o.value.pushNotify=c)},null,8,["modelValue"])]),_:1},8,["label"]),t(p,{field:"advertisingContent",label:e.$t("stepForm.form.label.advertisingContent"),rules:[{required:!0,message:e.$t("stepForm.form.error.advertisingContent.required")},{maxLength:200,message:e.$t("stepForm.form.error.advertisingContent.maxLength")}],"row-class":"keep-margin"},{default:a(()=>[t(k,{modelValue:o.value.advertisingContent,"onUpdate:modelValue":l[4]||(l[4]=c=>o.value.advertisingContent=c),placeholder:e.$t("stepForm.placeholder.advertisingContent")},null,8,["modelValue","placeholder"])]),_:1},8,["label","rules"]),t(p,null,{default:a(()=>[t(B,null,{default:a(()=>[t(u,{type:"secondary",onClick:r},{default:a(()=>[n(i(e.$t("stepForm.button.prev")),1)]),_:1}),t(u,{type:"primary",onClick:b},{default:a(()=>[n(i(e.$t("stepForm.button.next")),1)]),_:1})]),_:1})]),_:1})]),_:1},8,["model"])}}});const re=V(ae,[["__scopeId","data-v-d889eaf1"]]),le={class:"success-wrap"},se={class:"details-wrapper"},ne=w({__name:"success",emits:["changeStep"],setup(h,{emit:v}){const d=v,s=()=>{d("changeStep",1)};return(o,b)=>{const r=z,e=S,l=M,m=Y,p=j,_=G;return f(),U("div",le,[t(r,{status:"success",title:o.$t("stepForm.success.title"),subtitle:o.$t("stepForm.success.subTitle")},null,8,["title","subtitle"]),t(l,{size:16},{default:a(()=>[t(e,{key:"view",type:"primary"},{default:a(()=>[n(i(o.$t("stepForm.button.view")),1)]),_:1}),t(e,{key:"again",type:"secondary",onClick:s},{default:a(()=>[n(i(o.$t("stepForm.button.again")),1)]),_:1})]),_:1}),C("div",se,[t(m,{heading:6,style:{"margin-top":"0"}},{default:a(()=>[n(i(o.$t("stepForm.form.description.title")),1)]),_:1}),t(_,{style:{"margin-bottom":"0"}},{default:a(()=>[n(i(o.$t("stepForm.form.description.text"))+" ",1),t(p,{href:"link"},{default:a(()=>[n(i(o.$t("stepForm.button.view")),1)]),_:1})]),_:1})])])}}});const pe=V(ne,[["__scopeId","data-v-c9d0c124"]]),ie={class:"container"},me={class:"wrapper"},ue={name:"Step"},de=w({...ue,setup(h){const{loading:v,setLoading:d}=x(!1),s=y(1),o=y({}),b=async()=>{d(!0);try{await ee(o.value),s.value=3,o.value={}}catch{}finally{d(!1)}},r=(e,l)=>{if(typeof e=="number"){s.value=e;return}if(e==="forward"||e==="submit"){if(o.value={...o.value,...l},e==="submit"){b();return}s.value+=1}else e==="backward"&&(s.value-=1)};return(e,l)=>{const m=H("Breadcrumb"),p=W,_=Z,$=J,F=X;return f(),U("div",ie,[t(m,{items:["menu.form","menu.form.step"]}),t(F,{loading:Q(v),style:{width:"100%"}},{default:a(()=>[t($,{class:"general-card"},{title:a(()=>[n(i(e.$t("stepForm.step.title")),1)]),default:a(()=>[C("div",me,[t(_,{current:s.value,"onUpdate:current":l[0]||(l[0]=k=>s.value=k),style:{width:"580px"},"line-less":"",class:"steps"},{default:a(()=>[t(p,{description:e.$t("stepForm.step.subTitle.baseInfo")},{default:a(()=>[n(i(e.$t("stepForm.step.title.baseInfo")),1)]),_:1},8,["description"]),t(p,{description:e.$t("stepForm.step.subTitle.channel")},{default:a(()=>[n(i(e.$t("stepForm.step.title.channel")),1)]),_:1},8,["description"]),t(p,{description:e.$t("stepForm.step.subTitle.finish")},{default:a(()=>[n(i(e.$t("stepForm.step.title.finish")),1)]),_:1},8,["description"])]),_:1},8,["current"]),(f(),g(K,null,[s.value===1?(f(),g(oe,{key:0,onChangeStep:r})):s.value===2?(f(),g(re,{key:1,onChangeStep:r})):s.value===3?(f(),g(pe,{key:2,onChangeStep:r})):O("",!0)],1024))])]),_:1})]),_:1},8,["loading"])])}}});const Be=V(de,[["__scopeId","data-v-6601e207"]]);export{Be as default};
