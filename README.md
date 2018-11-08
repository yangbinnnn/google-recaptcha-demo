# google recaptcha demo

0. 从https://developers.google.com/recaptcha 获取`sitekey` 和`secret`
1. 前端页面加载通过`sitekey` 加载`g-recaptcha`
2. 用户与`g-recaptcha` 交互，谷歌实时对用户操作进行验证
3. 前端页面绑定特定事件获取用户验证`response`，推送到后端服务双重验证，后端服务将`response` + `secret` 推送到谷歌
4. 后端服务验证通过，则允许用户进行操作
