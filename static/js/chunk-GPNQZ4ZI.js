var k=(f,g,b)=>new Promise((h,e)=>{var i=a=>{try{c(b.next(a))}catch(d){e(d)}},j=a=>{try{c(b.throw(a))}catch(d){e(d)}},c=a=>a.done?h(a.value):Promise.resolve(a.value).then(i,j);c((b=b.apply(f,g)).next())});export{k as a};
//# sourceMappingURL=chunk-GPNQZ4ZI.js.map
