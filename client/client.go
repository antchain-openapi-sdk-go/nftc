// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

/**
 * Model for initing client
 */
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 通用资源差量包信息
type GeneralResourcePatch struct {
	// 差量包文件下载地址
	Url *string `json:"url,omitempty" xml:"url,omitempty" require:"true"`
	// 差量包的旧版本号
	OldVersion *int64 `json:"old_version,omitempty" xml:"old_version,omitempty" require:"true"`
	// md5摘要值
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty" require:"true"`
	// 差量包文件大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty" require:"true"`
}

func (s GeneralResourcePatch) String() string {
	return tea.Prettify(s)
}

func (s GeneralResourcePatch) GoString() string {
	return s.String()
}

func (s *GeneralResourcePatch) SetUrl(v string) *GeneralResourcePatch {
	s.Url = &v
	return s
}

func (s *GeneralResourcePatch) SetOldVersion(v int64) *GeneralResourcePatch {
	s.OldVersion = &v
	return s
}

func (s *GeneralResourcePatch) SetMd5(v string) *GeneralResourcePatch {
	s.Md5 = &v
	return s
}

func (s *GeneralResourcePatch) SetSize(v int64) *GeneralResourcePatch {
	s.Size = &v
	return s
}

type PublishMerchantDiyskuRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 业务请求id，用来做业务上的幂等。后面查询状态也是此字段
	BizNo *string `json:"biz_no,omitempty" xml:"biz_no,omitempty" require:"true"`
	// 调用渠道
	// PET 宠物
	// MEMBER 会员
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty" require:"true"`
	// 领取数字藏品的用户ID，支持2088/手机号/1322
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty" require:"true"`
	// 支付宝2088账号：ALIPAY_UID
	// 手机号：PHONE_NO
	// 鲸探1322账号：FANS_UID
	UserType *string `json:"user_type,omitempty" xml:"user_type,omitempty" require:"true"`
	// 需要发放的SKUID编码
	SkuId *string `json:"sku_id,omitempty" xml:"sku_id,omitempty" require:"true"`
	// 一期仅支持图片：IMAGE
	SkuType *string `json:"sku_type,omitempty" xml:"sku_type,omitempty" require:"true"`
	// 数字藏品的缩略图地址，可与原图相同也可不同，需要校验长宽比为1:1
	ThumbnailUrl *string `json:"thumbnail_url,omitempty" xml:"thumbnail_url,omitempty" require:"true"`
	// 数字藏品的原图地址，需要校验长宽比为1:1
	OriginalUrl *string `json:"original_url,omitempty" xml:"original_url,omitempty" require:"true"`
}

func (s PublishMerchantDiyskuRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishMerchantDiyskuRequest) GoString() string {
	return s.String()
}

func (s *PublishMerchantDiyskuRequest) SetAuthToken(v string) *PublishMerchantDiyskuRequest {
	s.AuthToken = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetProductInstanceId(v string) *PublishMerchantDiyskuRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetBizNo(v string) *PublishMerchantDiyskuRequest {
	s.BizNo = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetChannel(v string) *PublishMerchantDiyskuRequest {
	s.Channel = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetUserId(v string) *PublishMerchantDiyskuRequest {
	s.UserId = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetUserType(v string) *PublishMerchantDiyskuRequest {
	s.UserType = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetSkuId(v string) *PublishMerchantDiyskuRequest {
	s.SkuId = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetSkuType(v string) *PublishMerchantDiyskuRequest {
	s.SkuType = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetThumbnailUrl(v string) *PublishMerchantDiyskuRequest {
	s.ThumbnailUrl = &v
	return s
}

func (s *PublishMerchantDiyskuRequest) SetOriginalUrl(v string) *PublishMerchantDiyskuRequest {
	s.OriginalUrl = &v
	return s
}

type PublishMerchantDiyskuResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 生成的nftid结果
	NftId *string `json:"nft_id,omitempty" xml:"nft_id,omitempty"`
	// 获得时间
	ReceiveTime *string `json:"receive_time,omitempty" xml:"receive_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 发放状态
	// AUDIT_SUBMIT("审核中"),
	// AUDIT_SUCCESS("审核通过"),
	// AUDIT_FAIL("审核未通过"),
	// TRANSFER_UNKNOWN("发放结果未知"),
	// TRANSFER_SUCCESS("发放成功"),
	// TRANSFER_FAILED("发放失败"),
	// SHIELD("屏蔽"),
	// RECYCLE("回收"),
	SendStatus *string `json:"send_status,omitempty" xml:"send_status,omitempty"`
}

func (s PublishMerchantDiyskuResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishMerchantDiyskuResponse) GoString() string {
	return s.String()
}

func (s *PublishMerchantDiyskuResponse) SetReqMsgId(v string) *PublishMerchantDiyskuResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PublishMerchantDiyskuResponse) SetResultCode(v string) *PublishMerchantDiyskuResponse {
	s.ResultCode = &v
	return s
}

func (s *PublishMerchantDiyskuResponse) SetResultMsg(v string) *PublishMerchantDiyskuResponse {
	s.ResultMsg = &v
	return s
}

func (s *PublishMerchantDiyskuResponse) SetNftId(v string) *PublishMerchantDiyskuResponse {
	s.NftId = &v
	return s
}

func (s *PublishMerchantDiyskuResponse) SetReceiveTime(v string) *PublishMerchantDiyskuResponse {
	s.ReceiveTime = &v
	return s
}

func (s *PublishMerchantDiyskuResponse) SetSendStatus(v string) *PublishMerchantDiyskuResponse {
	s.SendStatus = &v
	return s
}

type QueryMerchantDiyskuRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 业务请求id，用来做业务上的幂等
	BizNo *string `json:"biz_no,omitempty" xml:"biz_no,omitempty" require:"true"`
	// 领取数字藏品的用户ID，支持2088/手机号/1322
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty" require:"true"`
}

func (s QueryMerchantDiyskuRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMerchantDiyskuRequest) GoString() string {
	return s.String()
}

func (s *QueryMerchantDiyskuRequest) SetAuthToken(v string) *QueryMerchantDiyskuRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryMerchantDiyskuRequest) SetProductInstanceId(v string) *QueryMerchantDiyskuRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryMerchantDiyskuRequest) SetBizNo(v string) *QueryMerchantDiyskuRequest {
	s.BizNo = &v
	return s
}

func (s *QueryMerchantDiyskuRequest) SetUserId(v string) *QueryMerchantDiyskuRequest {
	s.UserId = &v
	return s
}

type QueryMerchantDiyskuResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 生成的nftid结果
	NftId *string `json:"nft_id,omitempty" xml:"nft_id,omitempty"`
	// 获得时间
	ReceiveTime *string `json:"receive_time,omitempty" xml:"receive_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 发放状态
	// AUDIT_SUBMIT("审核中"),
	// AUDIT_SUCCESS("审核通过"),
	// AUDIT_FAIL("审核未通过"),
	// TRANSFER_UNKNOWN("发放结果未知"),
	// TRANSFER_SUCCESS("发放成功"),
	// TRANSFER_FAILED("发放失败"),
	// SHIELD("屏蔽"),
	// RECYCLE("回收"),
	SendStatus *string `json:"send_status,omitempty" xml:"send_status,omitempty"`
}

func (s QueryMerchantDiyskuResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMerchantDiyskuResponse) GoString() string {
	return s.String()
}

func (s *QueryMerchantDiyskuResponse) SetReqMsgId(v string) *QueryMerchantDiyskuResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryMerchantDiyskuResponse) SetResultCode(v string) *QueryMerchantDiyskuResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryMerchantDiyskuResponse) SetResultMsg(v string) *QueryMerchantDiyskuResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryMerchantDiyskuResponse) SetNftId(v string) *QueryMerchantDiyskuResponse {
	s.NftId = &v
	return s
}

func (s *QueryMerchantDiyskuResponse) SetReceiveTime(v string) *QueryMerchantDiyskuResponse {
	s.ReceiveTime = &v
	return s
}

func (s *QueryMerchantDiyskuResponse) SetSendStatus(v string) *QueryMerchantDiyskuResponse {
	s.SendStatus = &v
	return s
}

type QueryMerchantUgcimagesRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// ugc资产铸造记录id列表
	RecordIdList []*string `json:"record_id_list,omitempty" xml:"record_id_list,omitempty" require:"true" type:"Repeated"`
	// 场景
	BizScene *string `json:"biz_scene,omitempty" xml:"biz_scene,omitempty" require:"true"`
}

func (s QueryMerchantUgcimagesRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryMerchantUgcimagesRequest) GoString() string {
	return s.String()
}

func (s *QueryMerchantUgcimagesRequest) SetAuthToken(v string) *QueryMerchantUgcimagesRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryMerchantUgcimagesRequest) SetProductInstanceId(v string) *QueryMerchantUgcimagesRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryMerchantUgcimagesRequest) SetRecordIdList(v []*string) *QueryMerchantUgcimagesRequest {
	s.RecordIdList = v
	return s
}

func (s *QueryMerchantUgcimagesRequest) SetBizScene(v string) *QueryMerchantUgcimagesRequest {
	s.BizScene = &v
	return s
}

type QueryMerchantUgcimagesResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 入参中id对应的图片列表
	ImgList []*string `json:"img_list,omitempty" xml:"img_list,omitempty" type:"Repeated"`
}

func (s QueryMerchantUgcimagesResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryMerchantUgcimagesResponse) GoString() string {
	return s.String()
}

func (s *QueryMerchantUgcimagesResponse) SetReqMsgId(v string) *QueryMerchantUgcimagesResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryMerchantUgcimagesResponse) SetResultCode(v string) *QueryMerchantUgcimagesResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryMerchantUgcimagesResponse) SetResultMsg(v string) *QueryMerchantUgcimagesResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryMerchantUgcimagesResponse) SetImgList(v []*string) *QueryMerchantUgcimagesResponse {
	s.ImgList = v
	return s
}

type CreateResourceGeneralresourceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 小程序的APP ID
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	// 业务幂等ID
	BizId *string `json:"biz_id,omitempty" xml:"biz_id,omitempty" require:"true"`
	// 是否需要生成差量包
	NeedPatch *bool `json:"need_patch,omitempty" xml:"need_patch,omitempty" require:"true"`
}

func (s CreateResourceGeneralresourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGeneralresourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceGeneralresourceRequest) SetAuthToken(v string) *CreateResourceGeneralresourceRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateResourceGeneralresourceRequest) SetProductInstanceId(v string) *CreateResourceGeneralresourceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateResourceGeneralresourceRequest) SetAppId(v string) *CreateResourceGeneralresourceRequest {
	s.AppId = &v
	return s
}

func (s *CreateResourceGeneralresourceRequest) SetBizId(v string) *CreateResourceGeneralresourceRequest {
	s.BizId = &v
	return s
}

func (s *CreateResourceGeneralresourceRequest) SetNeedPatch(v bool) *CreateResourceGeneralresourceRequest {
	s.NeedPatch = &v
	return s
}

type CreateResourceGeneralresourceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 文件上传地址
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 资源上传的授权token，上传文件时需要将该值放入header中进行鉴权
	MaasToken *string `json:"maas_token,omitempty" xml:"maas_token,omitempty"`
}

func (s CreateResourceGeneralresourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceGeneralresourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceGeneralresourceResponse) SetReqMsgId(v string) *CreateResourceGeneralresourceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateResourceGeneralresourceResponse) SetResultCode(v string) *CreateResourceGeneralresourceResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateResourceGeneralresourceResponse) SetResultMsg(v string) *CreateResourceGeneralresourceResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateResourceGeneralresourceResponse) SetResourceId(v string) *CreateResourceGeneralresourceResponse {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceGeneralresourceResponse) SetUrl(v string) *CreateResourceGeneralresourceResponse {
	s.Url = &v
	return s
}

func (s *CreateResourceGeneralresourceResponse) SetMaasToken(v string) *CreateResourceGeneralresourceResponse {
	s.MaasToken = &v
	return s
}

type BindResourceGeneralresourcefileRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 小程序的APP ID
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty" require:"true"`
	// 文件ID
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty" require:"true"`
}

func (s BindResourceGeneralresourcefileRequest) String() string {
	return tea.Prettify(s)
}

func (s BindResourceGeneralresourcefileRequest) GoString() string {
	return s.String()
}

func (s *BindResourceGeneralresourcefileRequest) SetAuthToken(v string) *BindResourceGeneralresourcefileRequest {
	s.AuthToken = &v
	return s
}

func (s *BindResourceGeneralresourcefileRequest) SetProductInstanceId(v string) *BindResourceGeneralresourcefileRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *BindResourceGeneralresourcefileRequest) SetAppId(v string) *BindResourceGeneralresourcefileRequest {
	s.AppId = &v
	return s
}

func (s *BindResourceGeneralresourcefileRequest) SetResourceId(v string) *BindResourceGeneralresourcefileRequest {
	s.ResourceId = &v
	return s
}

func (s *BindResourceGeneralresourcefileRequest) SetFileId(v string) *BindResourceGeneralresourcefileRequest {
	s.FileId = &v
	return s
}

type BindResourceGeneralresourcefileResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 版本号
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 文件下载地址
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// md5摘要值
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 文件大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s BindResourceGeneralresourcefileResponse) String() string {
	return tea.Prettify(s)
}

func (s BindResourceGeneralresourcefileResponse) GoString() string {
	return s.String()
}

func (s *BindResourceGeneralresourcefileResponse) SetReqMsgId(v string) *BindResourceGeneralresourcefileResponse {
	s.ReqMsgId = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetResultCode(v string) *BindResourceGeneralresourcefileResponse {
	s.ResultCode = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetResultMsg(v string) *BindResourceGeneralresourcefileResponse {
	s.ResultMsg = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetVersion(v int64) *BindResourceGeneralresourcefileResponse {
	s.Version = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetUrl(v string) *BindResourceGeneralresourcefileResponse {
	s.Url = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetMd5(v string) *BindResourceGeneralresourcefileResponse {
	s.Md5 = &v
	return s
}

func (s *BindResourceGeneralresourcefileResponse) SetSize(v int64) *BindResourceGeneralresourcefileResponse {
	s.Size = &v
	return s
}

type PublishResourceGeneralresourcefileRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 小程序的APP ID
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	// 业务幂等ID
	BizId *string `json:"biz_id,omitempty" xml:"biz_id,omitempty" require:"true"`
	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty" require:"true"`
	// 版本号
	FileVersion *int64 `json:"file_version,omitempty" xml:"file_version,omitempty" require:"true"`
}

func (s PublishResourceGeneralresourcefileRequest) String() string {
	return tea.Prettify(s)
}

func (s PublishResourceGeneralresourcefileRequest) GoString() string {
	return s.String()
}

func (s *PublishResourceGeneralresourcefileRequest) SetAuthToken(v string) *PublishResourceGeneralresourcefileRequest {
	s.AuthToken = &v
	return s
}

func (s *PublishResourceGeneralresourcefileRequest) SetProductInstanceId(v string) *PublishResourceGeneralresourcefileRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PublishResourceGeneralresourcefileRequest) SetAppId(v string) *PublishResourceGeneralresourcefileRequest {
	s.AppId = &v
	return s
}

func (s *PublishResourceGeneralresourcefileRequest) SetBizId(v string) *PublishResourceGeneralresourcefileRequest {
	s.BizId = &v
	return s
}

func (s *PublishResourceGeneralresourcefileRequest) SetResourceId(v string) *PublishResourceGeneralresourcefileRequest {
	s.ResourceId = &v
	return s
}

func (s *PublishResourceGeneralresourcefileRequest) SetFileVersion(v int64) *PublishResourceGeneralresourcefileRequest {
	s.FileVersion = &v
	return s
}

type PublishResourceGeneralresourcefileResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 版本号
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 资源文件下载地址
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// md5摘要值
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 文件大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s PublishResourceGeneralresourcefileResponse) String() string {
	return tea.Prettify(s)
}

func (s PublishResourceGeneralresourcefileResponse) GoString() string {
	return s.String()
}

func (s *PublishResourceGeneralresourcefileResponse) SetReqMsgId(v string) *PublishResourceGeneralresourcefileResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetResultCode(v string) *PublishResourceGeneralresourcefileResponse {
	s.ResultCode = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetResultMsg(v string) *PublishResourceGeneralresourcefileResponse {
	s.ResultMsg = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetVersion(v int64) *PublishResourceGeneralresourcefileResponse {
	s.Version = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetUrl(v string) *PublishResourceGeneralresourcefileResponse {
	s.Url = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetMd5(v string) *PublishResourceGeneralresourcefileResponse {
	s.Md5 = &v
	return s
}

func (s *PublishResourceGeneralresourcefileResponse) SetSize(v int64) *PublishResourceGeneralresourcefileResponse {
	s.Size = &v
	return s
}

type QueryResourcePatchlistRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 小程序的APP ID
	AppId *string `json:"app_id,omitempty" xml:"app_id,omitempty" require:"true"`
	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty" require:"true"`
	// 版本号
	FileVersion *int64 `json:"file_version,omitempty" xml:"file_version,omitempty" require:"true"`
}

func (s QueryResourcePatchlistRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePatchlistRequest) GoString() string {
	return s.String()
}

func (s *QueryResourcePatchlistRequest) SetAuthToken(v string) *QueryResourcePatchlistRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryResourcePatchlistRequest) SetProductInstanceId(v string) *QueryResourcePatchlistRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryResourcePatchlistRequest) SetAppId(v string) *QueryResourcePatchlistRequest {
	s.AppId = &v
	return s
}

func (s *QueryResourcePatchlistRequest) SetResourceId(v string) *QueryResourcePatchlistRequest {
	s.ResourceId = &v
	return s
}

func (s *QueryResourcePatchlistRequest) SetFileVersion(v int64) *QueryResourcePatchlistRequest {
	s.FileVersion = &v
	return s
}

type QueryResourcePatchlistResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 差量包生成状态，包括PATCHING（生成中）和FINISHED（生成完成）
	PatchStatus *string `json:"patch_status,omitempty" xml:"patch_status,omitempty"`
	// 全量包下载地址
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// 资源ID
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id,omitempty"`
	// 全量包版本号
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 全量包md5摘要值
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// 全量包文件大小
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 差量包文件信息列表
	PatchList []*GeneralResourcePatch `json:"patch_list,omitempty" xml:"patch_list,omitempty" type:"Repeated"`
}

func (s QueryResourcePatchlistResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryResourcePatchlistResponse) GoString() string {
	return s.String()
}

func (s *QueryResourcePatchlistResponse) SetReqMsgId(v string) *QueryResourcePatchlistResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetResultCode(v string) *QueryResourcePatchlistResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetResultMsg(v string) *QueryResourcePatchlistResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetPatchStatus(v string) *QueryResourcePatchlistResponse {
	s.PatchStatus = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetUrl(v string) *QueryResourcePatchlistResponse {
	s.Url = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetResourceId(v string) *QueryResourcePatchlistResponse {
	s.ResourceId = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetVersion(v int64) *QueryResourcePatchlistResponse {
	s.Version = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetMd5(v string) *QueryResourcePatchlistResponse {
	s.Md5 = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetSize(v int64) *QueryResourcePatchlistResponse {
	s.Size = &v
	return s
}

func (s *QueryResourcePatchlistResponse) SetPatchList(v []*GeneralResourcePatch) *QueryResourcePatchlistResponse {
	s.PatchList = v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

/**
 * Init client with Config
 * @param config config contains the necessary information to create a client
 */
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param protocol http or https
 * @param method e.g. GET
 * @param pathname pathname of every api
 * @param request which contains request params
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":          "retry",
		"readTimeout":        tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":     tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":          tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":         tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":            tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":       tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":  tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDuration":  tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":        tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost": tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.0.4"),
				"_prod_code":       tea.String("NFTC"),
				"_prod_channel":    tea.String("undefined"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res, _err := util.AssertAsMap(obj)
			if _err != nil {
				return _result, _err
			}

			resp, _err := util.AssertAsMap(res["response"])
			if _err != nil {
				return _result, _err
			}

			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Description: DIY藏品发放
 * Summary: DIY藏品发放
 */
func (client *Client) PublishMerchantDiysku(request *PublishMerchantDiyskuRequest) (_result *PublishMerchantDiyskuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishMerchantDiyskuResponse{}
	_body, _err := client.PublishMerchantDiyskuEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: DIY藏品发放
 * Summary: DIY藏品发放
 */
func (client *Client) PublishMerchantDiyskuEx(request *PublishMerchantDiyskuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishMerchantDiyskuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PublishMerchantDiyskuResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.merchant.diysku.publish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: DIY藏品发放查询
 * Summary: DIY藏品发放查询
 */
func (client *Client) QueryMerchantDiysku(request *QueryMerchantDiyskuRequest) (_result *QueryMerchantDiyskuResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMerchantDiyskuResponse{}
	_body, _err := client.QueryMerchantDiyskuEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: DIY藏品发放查询
 * Summary: DIY藏品发放查询
 */
func (client *Client) QueryMerchantDiyskuEx(request *QueryMerchantDiyskuRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryMerchantDiyskuResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryMerchantDiyskuResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.merchant.diysku.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 提供ISV  ugc铸造图片查询的openapi服务
 * Summary: ugc铸造图片查询的openapi接口
 */
func (client *Client) QueryMerchantUgcimages(request *QueryMerchantUgcimagesRequest) (_result *QueryMerchantUgcimagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryMerchantUgcimagesResponse{}
	_body, _err := client.QueryMerchantUgcimagesEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 提供ISV  ugc铸造图片查询的openapi服务
 * Summary: ugc铸造图片查询的openapi接口
 */
func (client *Client) QueryMerchantUgcimagesEx(request *QueryMerchantUgcimagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryMerchantUgcimagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryMerchantUgcimagesResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.merchant.ugcimages.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建通用资源
 * Summary: 创建通用资源
 */
func (client *Client) CreateResourceGeneralresource(request *CreateResourceGeneralresourceRequest) (_result *CreateResourceGeneralresourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceGeneralresourceResponse{}
	_body, _err := client.CreateResourceGeneralresourceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建通用资源
 * Summary: 创建通用资源
 */
func (client *Client) CreateResourceGeneralresourceEx(request *CreateResourceGeneralresourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceGeneralresourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateResourceGeneralresourceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.resource.generalresource.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 绑定文件到通用资源ID
 * Summary: 绑定文件到通用资源ID
 */
func (client *Client) BindResourceGeneralresourcefile(request *BindResourceGeneralresourcefileRequest) (_result *BindResourceGeneralresourcefileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindResourceGeneralresourcefileResponse{}
	_body, _err := client.BindResourceGeneralresourcefileEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 绑定文件到通用资源ID
 * Summary: 绑定文件到通用资源ID
 */
func (client *Client) BindResourceGeneralresourcefileEx(request *BindResourceGeneralresourcefileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindResourceGeneralresourcefileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BindResourceGeneralresourcefileResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.resource.generalresourcefile.bind"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 发布通用资源文件
 * Summary: 发布通用资源文件
 */
func (client *Client) PublishResourceGeneralresourcefile(request *PublishResourceGeneralresourcefileRequest) (_result *PublishResourceGeneralresourcefileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PublishResourceGeneralresourcefileResponse{}
	_body, _err := client.PublishResourceGeneralresourcefileEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 发布通用资源文件
 * Summary: 发布通用资源文件
 */
func (client *Client) PublishResourceGeneralresourcefileEx(request *PublishResourceGeneralresourcefileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PublishResourceGeneralresourcefileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PublishResourceGeneralresourcefileResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.resource.generalresourcefile.publish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询资源文件差量包列表
 * Summary: 查询资源文件差量包列表
 */
func (client *Client) QueryResourcePatchlist(request *QueryResourcePatchlistRequest) (_result *QueryResourcePatchlistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryResourcePatchlistResponse{}
	_body, _err := client.QueryResourcePatchlistEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询资源文件差量包列表
 * Summary: 查询资源文件差量包列表
 */
func (client *Client) QueryResourcePatchlistEx(request *QueryResourcePatchlistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryResourcePatchlistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryResourcePatchlistResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("antchain.nftc.resource.patchlist.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}