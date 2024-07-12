"use strict";(self.webpackChunkweb=self.webpackChunkweb||[]).push([[849],{8849:(b,S,n)=>{n.r(S),n.d(S,{DashboardModule:()=>f});var u=n(8583),r=n(9502),_=n(476),g=n(2468),v=n(2437),A=n(8391),D=n(9744),l=n(9386),C=n(3692);let O=(()=>{class i{constructor(e,t,s,a,c,E){this.dataService=e,this.metadataSvc=t,this.schedulerSvc=s,this.notiSvc=a,this.systemAgentSvc=c,this.registrySvc=E,this.eventCount=0,this.readingCount=0,this.deviceSvcCount=0,this.deviceSvcStatusLockedCount=0,this.deviceCount=0,this.deviceStatusLockedCount=0,this.deviceProfileCount=0,this.schedulerCount=0,this.notificationCount=0,this.registeredServiceCount=0}ngOnInit(){this.dataService.ping().subscribe(()=>{this.getEventAndReadingCount()}),this.metadataSvc.ping().subscribe(()=>{this.getDeviceServiceCount(),this.getDeviceCount(),this.getDeviceProfileCount()}),this.schedulerSvc.ping().subscribe(()=>{this.getIntervalCount()}),this.notiSvc.ping().subscribe(()=>{this.getNotificationCount()}),this.systemAgentSvc.ping().subscribe(()=>{this.registrySvc.ping().subscribe(()=>{this.getRegisteredServiceCount()})})}getEventAndReadingCount(){this.dataService.eventCount().subscribe(e=>this.eventCount=e.Count),this.dataService.readingCount().subscribe(e=>this.readingCount=e.Count)}getDeviceServiceCount(){this.metadataSvc.allDeviceServices().subscribe(e=>{this.deviceSvcCount=e.services.length,e.services.forEach((t,s)=>{"LOCKED"===t.adminState&&this.deviceSvcStatusLockedCount++})})}getDeviceCount(){this.metadataSvc.allDevices().subscribe(e=>{this.deviceCount=e.devices.length,e.devices.forEach((t,s)=>{"LOCKED"===t.adminState&&this.deviceStatusLockedCount++})})}getDeviceProfileCount(){this.metadataSvc.allDeviceProfolesPagination(0,-1).subscribe(e=>{this.deviceProfileCount=e.profiles.length})}getIntervalCount(){this.schedulerSvc.findAllIntervalsPagination(0,-1).subscribe(e=>{this.schedulerCount=e.intervals.length})}getNotificationCount(){this.notiSvc.findNotificationsByStatusPagination(0,-1,"NEW").subscribe(e=>{this.notificationCount=e.notifications.length})}getRegisteredServiceCount(){this.systemAgentSvc.getRegisteredServiceAll().subscribe(e=>{this.registeredServiceCount=e.length?e.length:0})}}return i.\u0275fac=function(e){return new(e||i)(_.Y36(g.D),_.Y36(v.D),_.Y36(A.G),_.Y36(D.T),_.Y36(l.J),_.Y36(C.r))},i.\u0275cmp=_.Xpm({type:i,selectors:[["app-dashboard"]],decls:67,vars:11,consts:function(){let o,e,t,s,a,c,E;return o=" Device Services ",e="\u8BBE\u5907",t="\u8BBE\u5907\u5143\u4FE1\u606F",s="\u4EFB\u52A1\u8C03\u5EA6\u4E2D\u5FC3",a="\u901A\u77E5\u670D\u52A1",c="\u4E8B\u4EF6\u603B\u6570",E="\u8BFB\u503C\u603B\u6570",[[1,"row"],[1,"col-lg-4"],["role","button","routerLink","/metadata",1,"card"],[1,"card-body"],[1,"card-title"],o,[1,"d-inline"],[1,"badge","badge-info"],[1,"float-right","badge","badge-danger"],[1,"float-right","badge","badge-success","mr-2"],["role","button","routerLink","/metadata/device-center",1,"card"],e,["role","button","routerLink","/metadata/device-profile-center",1,"card"],t,[1,"row","mt-3"],[1,"col-lg-6"],["role","button","routerLink","/scheduler",1,"card"],s,["href","#",1,"card-link","font-weight-bolder","badge","badge-info"],["role","button","routerLink","/notifications",1,"card"],a,["role","button","routerLink","/core-data",1,"card"],c,E]},template:function(e,t){1&e&&(_.TgZ(0,"div",0),_.TgZ(1,"div",1),_.TgZ(2,"div",2),_.TgZ(3,"div",3),_.TgZ(4,"h5",4),_.SDv(5,5),_.qZA(),_.TgZ(6,"h5",6),_.TgZ(7,"span",7),_._uU(8),_.qZA(),_.TgZ(9,"span",8),_._uU(10),_.qZA(),_.TgZ(11,"span",9),_._uU(12),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(13,"div",1),_.TgZ(14,"div",10),_.TgZ(15,"div",3),_.TgZ(16,"h5",4),_.SDv(17,11),_.qZA(),_.TgZ(18,"h5",6),_.TgZ(19,"span",7),_._uU(20),_.qZA(),_.TgZ(21,"span",8),_._uU(22),_.qZA(),_.TgZ(23,"span",9),_._uU(24),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(25,"div",1),_.TgZ(26,"div",12),_.TgZ(27,"div",3),_.TgZ(28,"h5",4),_.SDv(29,13),_.qZA(),_.TgZ(30,"h5",6),_.TgZ(31,"span",7),_._uU(32),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(33,"div",14),_.TgZ(34,"div",15),_.TgZ(35,"div",16),_.TgZ(36,"div",3),_.TgZ(37,"h5",4),_.SDv(38,17),_.qZA(),_.TgZ(39,"h5"),_.TgZ(40,"a",18),_._uU(41),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(42,"div",15),_.TgZ(43,"div",19),_.TgZ(44,"div",3),_.TgZ(45,"h5",4),_.SDv(46,20),_.qZA(),_.TgZ(47,"h5"),_.TgZ(48,"a",18),_._uU(49),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(50,"div",14),_.TgZ(51,"div",15),_.TgZ(52,"div",21),_.TgZ(53,"div",3),_.TgZ(54,"h5",4),_.SDv(55,22),_.qZA(),_.TgZ(56,"h5"),_.TgZ(57,"a",18),_._uU(58),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.TgZ(59,"div",15),_.TgZ(60,"div",21),_.TgZ(61,"div",3),_.TgZ(62,"h5",4),_.SDv(63,23),_.qZA(),_.TgZ(64,"h5"),_.TgZ(65,"a",18),_._uU(66),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA(),_.qZA()),2&e&&(_.xp6(8),_.hij(" ",t.deviceSvcCount>100?"100+":t.deviceSvcCount," "),_.xp6(2),_.hij("Locked ",t.deviceSvcStatusLockedCount,""),_.xp6(2),_.hij("Unlocked ",t.deviceSvcCount-t.deviceSvcStatusLockedCount,""),_.xp6(8),_.hij(" ",t.deviceCount>100?"100+":t.deviceCount," "),_.xp6(2),_.hij("Locked ",t.deviceStatusLockedCount,""),_.xp6(2),_.hij("Unlocked ",t.deviceCount-t.deviceStatusLockedCount,""),_.xp6(8),_.Oqu(t.deviceProfileCount>100?"100+":t.deviceProfileCount),_.xp6(9),_.Oqu(t.schedulerCount>100?"100+":t.schedulerCount),_.xp6(8),_.Oqu(t.notificationCount>100?"100+":t.notificationCount),_.xp6(9),_.Oqu(t.eventCount),_.xp6(8),_.Oqu(t.readingCount))},directives:[r.rH],styles:[".shadow[_ngcontent-%COMP%]{box-shadow:0 .5rem 1rem #00000026!important;border-radius:.25rem!important}.card[_ngcontent-%COMP%]:hover{box-shadow:0 .5rem 1rem #00000026!important;border-radius:.25rem!important}"]}),i})();const T=[{path:"",canActivate:[n(8253).a],component:O}];let Z=(()=>{class i{}return i.\u0275fac=function(e){return new(e||i)},i.\u0275mod=_.oAB({type:i}),i.\u0275inj=_.cJS({imports:[[r.Bz.forChild(T)],r.Bz]}),i})(),f=(()=>{class i{}return i.\u0275fac=function(e){return new(e||i)},i.\u0275mod=_.oAB({type:i}),i.\u0275inj=_.cJS({imports:[[u.ez,Z]]}),i})()}}]);