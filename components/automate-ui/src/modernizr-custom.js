/*
Modernizer has you create a build script with custom options
Click on the link below to be taken to the current configuration
to add more options and then replace the script below.
*/

/*! modernizr 3.5.0 (Custom Build) | MIT *
 * https://modernizr.com/download/?-fontface-inlinesvg-inputtypes-smil-svgasimg-atrule-hasevent-prefixed-setclasses-testallprops !*/
 !function(e,t,n){function r(e,t){return typeof e===t}function i(){var e,t,n,i,o,s,a;for(var l in S)if(S.hasOwnProperty(l)){if(e=[],t=S[l],t.name&&(e.push(t.name.toLowerCase()),t.options&&t.options.aliases&&t.options.aliases.length))for(n=0;n<t.options.aliases.length;n++)e.push(t.options.aliases[n].toLowerCase());for(i=r(t.fn,"function")?t.fn():t.fn,o=0;o<e.length;o++)s=e[o],a=s.split("."),1===a.length?Modernizr[a[0]]=i:(!Modernizr[a[0]]||Modernizr[a[0]]instanceof Boolean||(Modernizr[a[0]]=new Boolean(Modernizr[a[0]])),Modernizr[a[0]][a[1]]=i),C.push((i?"":"no-")+a.join("-"))}}function o(e){var t=b.className,n=Modernizr._config.classPrefix||"";if(x&&(t=t.baseVal),Modernizr._config.enableJSClass){var r=new RegExp("(^|\\s)"+n+"no-js(\\s|$)");t=t.replace(r,"$1"+n+"js$2")}Modernizr._config.enableClasses&&(t+=" "+n+e.join(" "+n),x?b.className.baseVal=t:b.className=t)}function s(){return"function"!=typeof t.createElement?t.createElement(arguments[0]):x?t.createElementNS.call(t,"http://www.w3.org/2000/svg",arguments[0]):t.createElement.apply(t,arguments)}function a(e){return e.replace(/([a-z])-([a-z])/g,function(e,t,n){return t+n.toUpperCase()}).replace(/^-/,"")}function l(e,t){if("object"==typeof e)for(var n in e)V(e,n)&&l(n,e[n]);else{e=e.toLowerCase();var r=e.split("."),i=Modernizr[r[0]];if(2==r.length&&(i=i[r[1]]),"undefined"!=typeof i)return Modernizr;t="function"==typeof t?t():t,1==r.length?Modernizr[r[0]]=t:(!Modernizr[r[0]]||Modernizr[r[0]]instanceof Boolean||(Modernizr[r[0]]=new Boolean(Modernizr[r[0]])),Modernizr[r[0]][r[1]]=t),o([(t&&0!=t?"":"no-")+r.join("-")]),Modernizr._trigger(e,t)}return Modernizr}function u(){var e=t.body;return e||(e=s(x?"svg":"body"),e.fake=!0),e}function f(e,n,r,i){var o,a,l,f,c="modernizr",d=s("div"),p=u();if(parseInt(r,10))for(;r--;)l=s("div"),l.id=i?i[r]:c+(r+1),d.appendChild(l);return o=s("style"),o.type="text/css",o.id="s"+c,(p.fake?p:d).appendChild(o),p.appendChild(d),o.styleSheet?o.styleSheet.cssText=e:o.appendChild(t.createTextNode(e)),d.id=c,p.fake&&(p.style.background="",p.style.overflow="hidden",f=b.style.overflow,b.style.overflow="hidden",b.appendChild(p)),a=n(d,e),p.fake?(p.parentNode.removeChild(p),b.style.overflow=f,b.offsetHeight):d.parentNode.removeChild(d),!!a}function c(e,t){return!!~(""+e).indexOf(t)}function d(e,t){return function(){return e.apply(t,arguments)}}function p(e,t,n){var i;for(var o in e)if(e[o]in t)return n===!1?e[o]:(i=t[e[o]],r(i,"function")?d(i,n||t):i);return!1}function m(e){return e.replace(/([A-Z])/g,function(e,t){return"-"+t.toLowerCase()}).replace(/^ms-/,"-ms-")}function h(t,n,r){var i;if("getComputedStyle"in e){i=getComputedStyle.call(e,t,n);var o=e.console;if(null!==i)r&&(i=i.getPropertyValue(r));else if(o){var s=o.error?"error":"log";o[s].call(o,"getComputedStyle returning null, its possible modernizr test results are inaccurate")}}else i=!n&&t.currentStyle&&t.currentStyle[r];return i}function g(t,r){var i=t.length;if("CSS"in e&&"supports"in e.CSS){for(;i--;)if(e.CSS.supports(m(t[i]),r))return!0;return!1}if("CSSSupportsRule"in e){for(var o=[];i--;)o.push("("+m(t[i])+":"+r+")");return o=o.join(" or "),f("@supports ("+o+") { #modernizr { position: absolute; } }",function(e){return"absolute"==h(e,null,"position")})}return n}function v(e,t,i,o){function l(){f&&(delete $.style,delete $.modElem)}if(o=r(o,"undefined")?!1:o,!r(i,"undefined")){var u=g(e,i);if(!r(u,"undefined"))return u}for(var f,d,p,m,h,v=["modernizr","tspan","samp"];!$.style&&v.length;)f=!0,$.modElem=s(v.shift()),$.style=$.modElem.style;for(p=e.length,d=0;p>d;d++)if(m=e[d],h=$.style[m],c(m,"-")&&(m=a(m)),$.style[m]!==n){if(o||r(i,"undefined"))return l(),"pfx"==t?m:!0;try{$.style[m]=i}catch(y){}if($.style[m]!=h)return l(),"pfx"==t?m:!0}return l(),!1}function y(e,t,n,i,o){var s=e.charAt(0).toUpperCase()+e.slice(1),a=(e+" "+k.join(s+" ")+s).split(" ");return r(t,"string")||r(t,"undefined")?v(a,t,i,o):(a=(e+" "+N.join(s+" ")+s).split(" "),p(a,t,n))}function w(e,t,r){return y(e,n,n,t,r)}var C=[],S=[],_={_version:"3.5.0",_config:{classPrefix:"",enableClasses:!0,enableJSClass:!0,usePrefixes:!0},_q:[],on:function(e,t){var n=this;setTimeout(function(){t(n[e])},0)},addTest:function(e,t,n){S.push({name:e,fn:t,options:n})},addAsyncTest:function(e){S.push({name:null,fn:e})}},Modernizr=function(){};Modernizr.prototype=_,Modernizr=new Modernizr;var b=t.documentElement,x="svg"===b.nodeName.toLowerCase(),T=function(){function e(e,t){var i;return e?(t&&"string"!=typeof t||(t=s(t||"div")),e="on"+e,i=e in t,!i&&r&&(t.setAttribute||(t=s("div")),t.setAttribute(e,""),i="function"==typeof t[e],t[e]!==n&&(t[e]=n),t.removeAttribute(e)),i):!1}var r=!("onblur"in t.documentElement);return e}();_.hasEvent=T,Modernizr.addTest("inlinesvg",function(){var e=s("div");return e.innerHTML="<svg/>","http://www.w3.org/2000/svg"==("undefined"!=typeof SVGRect&&e.firstChild&&e.firstChild.namespaceURI)});var E=s("input"),A="search tel url email datetime date month week time datetime-local number range color".split(" "),P={};Modernizr.inputtypes=function(e){for(var r,i,o,s=e.length,a="1)",l=0;s>l;l++)E.setAttribute("type",r=e[l]),o="text"!==E.type&&"style"in E,o&&(E.value=a,E.style.cssText="position:absolute;visibility:hidden;",/^range$/.test(r)&&E.style.WebkitAppearance!==n?(b.appendChild(E),i=t.defaultView,o=i.getComputedStyle&&"textfield"!==i.getComputedStyle(E,null).WebkitAppearance&&0!==E.offsetHeight,b.removeChild(E)):/^(search|tel)$/.test(r)||(o=/^(url|email)$/.test(r)?E.checkValidity&&E.checkValidity()===!1:E.value!=a)),P[e[l]]=!!o;return P}(A);var R="Moz O ms Webkit",k=_._config.usePrefixes?R.split(" "):[];_._cssomPrefixes=k;var z=function(t){var r,i=prefixes.length,o=e.CSSRule;if("undefined"==typeof o)return n;if(!t)return!1;if(t=t.replace(/^@/,""),r=t.replace(/-/g,"_").toUpperCase()+"_RULE",r in o)return"@"+t;for(var s=0;i>s;s++){var a=prefixes[s],l=a.toUpperCase()+"_"+r;if(l in o)return"@-"+a.toLowerCase()+"-"+t}return!1};_.atRule=z;var N=_._config.usePrefixes?R.toLowerCase().split(" "):[];_._domPrefixes=N;var V;!function(){var e={}.hasOwnProperty;V=r(e,"undefined")||r(e.call,"undefined")?function(e,t){return t in e&&r(e.constructor.prototype[t],"undefined")}:function(t,n){return e.call(t,n)}}(),_._l={},_.on=function(e,t){this._l[e]||(this._l[e]=[]),this._l[e].push(t),Modernizr.hasOwnProperty(e)&&setTimeout(function(){Modernizr._trigger(e,Modernizr[e])},0)},_._trigger=function(e,t){if(this._l[e]){var n=this._l[e];setTimeout(function(){var e,r;for(e=0;e<n.length;e++)(r=n[e])(t)},0),delete this._l[e]}},Modernizr._q.push(function(){_.addTest=l}),Modernizr.addTest("svgasimg",t.implementation.hasFeature("http://www.w3.org/TR/SVG11/feature#Image","1.1"));var j=_.testStyles=f,L=function(){var e=navigator.userAgent,t=e.match(/w(eb)?osbrowser/gi),n=e.match(/windows phone/gi)&&e.match(/iemobile\/([0-9])+/gi)&&parseFloat(RegExp.$1)>=9;return t||n}();L?Modernizr.addTest("fontface",!1):j('@font-face {font-family:"font";src:url("https://")}',function(e,n){var r=t.getElementById("smodernizr"),i=r.sheet||r.styleSheet,o=i?i.cssRules&&i.cssRules[0]?i.cssRules[0].cssText:i.cssText||"":"",s=/src/i.test(o)&&0===o.indexOf(n.split(" ")[0]);Modernizr.addTest("fontface",s)});var O={elem:s("modernizr")};Modernizr._q.push(function(){delete O.elem});var $={style:O.elem.style};Modernizr._q.unshift(function(){delete $.style}),_.testAllProps=y;_.prefixed=function(e,t,n){return 0===e.indexOf("@")?z(e):(-1!=e.indexOf("-")&&(e=a(e)),t?y(e,t,n):y(e,"pfx"))};_.testAllProps=w;var q={}.toString;Modernizr.addTest("smil",function(){return!!t.createElementNS&&/SVGAnimate/.test(q.call(t.createElementNS("http://www.w3.org/2000/svg","animate")))}),i(),o(C),delete _.addTest,delete _.addAsyncTest;for(var U=0;U<Modernizr._q.length;U++)Modernizr._q[U]();e.Modernizr=Modernizr}(window,document);