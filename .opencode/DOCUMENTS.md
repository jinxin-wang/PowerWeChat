# 微信小店文档链接汇总

## 通用 接口列表

### 获取稳定版接口调用凭据
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getstableaccesstoken.html

### 查询API调用额度
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getapiquota.html

### 重置指定API调用次数
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearapiquota.html

### 重置API调用次数
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearquota.html

### 使用AppSecret重置API调用次数
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearquotabyappsecret.html

### 网络通信检测
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_callbackcheck.html

### 获取微信API服务器IP
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getapidomainip.html

### 获取微信推送服务器IP
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getcallbackip.html

### 查询rid信息
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getridinfo.html

### 通过mediaid获取数据
https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getdatabymediaid.html


## 店铺管理 接口列表

### 获取店铺基本信息
https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_mmecapi_basicinfo.html

### 获取店铺二维码
https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_getshopqrcode.html

### 获取店铺口令
通过该接口可以获取微信小店的店铺微信口令，支持传入企业微信参数。通过带企业微信参数的口令下单的订单会在成交来源展示该企业。
https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_getshoptaglink.html


## 订单管理 接口列表

### 获取订单列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getorderlist.html

### 获取订单详情
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getorder.html

### 订单搜索
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_searchorder.html

### 修改订单价格
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changeorderprice.html

### 修改订单备注
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changemerchantnotes.html

### 修改订单地址
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changeorderaddress.html

### 修改物流信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changedeliveryinfo.html

### 同意用户修改收货地址申请
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_acceptorderaddressmodifyapply.html

### 拒绝用户修改收货地址申请
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_rejectorderaddressmodifyapply.html

### 上传生鲜质检信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_submitfreshinspectinfo.html

### 礼物订单新增备注信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_presentnote.html

### 获取礼物单的子单列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getpresentsuborder.html

### 获取所有待发货前更换sku待处理请求
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getpreshipmentchangeskuwaithandlelist.html

### 同意待发货前更换sku请求
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_approvepreshipmentchangesku.html

### 拒绝待发货前更换sku请求
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_rejectpreshipmentchangesku.html

### 解密订单中的详细收货信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_decodesensitiveinfo.html

### 申请查看订单真实号码
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_applyrealnumber.html

### 查看订单真实号审核状态
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getrealnumberviewaudit.html

### 订单再次申请虚拟号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_applyvirtualnumberagain.html

### 订单虚拟号延期
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_delayvirtualnumber.html

### 添加待认证的手机号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_privatenumberaddphone.html

### 获取短信验证码
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_privatenumbersendverifycode.html

### 获取小店手机号认证状态
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_privatenumbergetshopphone.html


## 售后管理 接口列表

### 获取售后单列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalelist.html

### 获取售后单
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersaleorder.html

### 同意售后
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_acceptapply.html

### 换货发货
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_acceptexchangereship.html

### 代用户发起售后
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_genaftersaleorder.html

### 商家获取保障单列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_searchguaranteeorder.html

### 获取保障单详情
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getguaranteeorder.html

### 商家同意保障单申请
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantacceptguarantee.html

### 商家协商保障单
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantmodifyguarantee.html

### 商家举证保障单
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantproofguarantee.html

### 商家拒绝保障单申请
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantrefuseguarantee.html

### 商家协商
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantupdateaftersale.html

### 获取全量售后原因
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalereason.html

### 拒绝售后
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_rejectapply.html

### 换货拒绝发货
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_rejectexchangereship.html

### 获取拒绝售后原因
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalerejectreason.html

### 上传退款凭证
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_uploadrefundcertificate.html

### 代用户发起退差价
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_refundpricediff.html

### 售后单兑换虚拟号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_applyvirtualtelnum.html


## 商家客服 接口列表

### 上传多媒体资源
https://developers.weixin.qq.com/doc/store/shop/API/kf/api_cosupload.html

### 发送消息
https://developers.weixin.qq.com/doc/store/shop/API/kf/api_sendmsg.html


## 纠纷管理 接口列表

### 商家补充纠纷单留言
https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_addcomplaintmaterial.html

### 商家举证
https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_addcomplaintproof.html

### 获取纠纷单
https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_getcomplaintorder.html


## 物流发货 接口列表

### 地址管理

#### 添加地址
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/address/api_addaddress.html

#### 获取地址列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/address/api_getaddresslist.html

#### 获取地址详情
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/address/api_getaddress.html

#### 更新地址
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/address/api_updateaddress.html

#### 删除地址
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/address/api_deleteaddress.html

### 运费模板

#### 增加运费模版
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/merchant/api_addfreighttemplate.html

#### 查询运费模版
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/merchant/api_getfreighttemplatedetail.html

#### 获取运费模板列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/merchant/api_getfreighttemplatelist.html

#### 更新运费模版
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/merchant/api_updatefreighttemplate.html

### 电子面单

#### 获取面单标准模板
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_get_template_config.html

#### 新增面单模板
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_createtemplate.html

#### 删除面单模版
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_deltemplate.html

#### 更新面单模版
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_updatetemplate.html

#### 获取面单模板信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_gettemplate.html

#### 根据模板ID获取面单模板信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_gettemplatebyid.html

#### 查询开通的电子面单网点/账号信息
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_getacct.html

#### 查询开通的快递公司列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_getdeliverylist.html

#### 电子面单预取号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_precreateorder.html

#### 电子面单取号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_createorder.html

#### 电子面单子件追加
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_addsuborder.html

#### 电子面单取消下单
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_cancelorder.html

#### 查询面单详情
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_getorder.html

#### 获取打印报文
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_get_print_content.html

#### 打印成功通知
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_printorder.html

#### 批量打印通知
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/ewaybill/api_ewaybill_batchprintorder.html

### 发货

#### 订单发货
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/delivery/api_senddelivery.html

#### 订单补发货
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/delivery/api_delivery_compensation.html

#### 获取快递公司列表
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/delivery/api_getdeliverycompanylistnew.html

#### 获取快递公司列表-旧
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/delivery/api_getdeliverycompanylist.html

### 物流公司虚拟号码

#### 获取虚拟号码池
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/phonenumber/api_getprivatenumberpool.html

#### 根据运单号获取真实手机号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/phonenumber/api_getrealnumber.html

#### 根据运单号获取虚拟手机号
https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-delivery/phonenumber/api_getvirtualnumber.html


## 优选联盟 接口列表

### 达人操作

#### 新增达人
https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_addpromoter.html

#### 删除达人
https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_deletepromoter.html

#### 获取达人详情信息
https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoter.html

#### 获取商店达人列表
https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoterlist.html

#### 编辑达人
https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_updpromoter.html

### 商品操作

#### 批量新增联盟商品
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchadditem.html

#### 删除联盟商品
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_deleteitem.html

#### 获取联盟商品详情
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitem.html

#### 批量新增联盟机构推广
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchaddheadsupplieritem.html

#### 获取联盟商品推广列表
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitemlist.html

#### 更新联盟商品信息
https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_upditem.html


## 罗盘商家版 接口列表

### 获取授权视频号列表
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderauthorizationlist.html

### 获取带货达人列表
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderlist.html

### 获取带货数据概览
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderoverall.html

### 获取带货达人商品列表
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderproductlist.html

### 获取带货达人详情
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderproductoverall.html

### 获取店铺开播列表
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshoplivelist.html

### 获取电商数据概览
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopoverall.html

### 获取商品详细信息
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopproductdata.html

### 获取商品列表
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopproductlist.html

### 获取店铺人群数据
https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopsaleprofiledata.html


## 小店会员 接口列表

### 获取用户积分
https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getvipuserscore.html

### 获取用户信息
https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserinfo.html

### 获取用户列表
https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserlist.html

### 获取用户积分流水
https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserscoreflowrecord.html

