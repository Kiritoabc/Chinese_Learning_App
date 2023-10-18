# 开发文档



功能需求：

- 提供外国网友学习中文的网站 or app
- 文化交流（尽量以乐趣为主）
- 注册登录功能
- 添加/删除视频功能
- 添加视频分组



前端界面简单设计





数据库设计:Chinese_Learning_DB

表:teaching_video(中文教学视频）

字段：

- id(唯一标识)
- name(视频名称)
- icon(视频图标)
- uri(视频的url路径)
- group:(视频所属分组)
- type:(类型)

表:cultural_exchange_video

字段:

- id(唯一标识)
- name(视频名称)
- icon(视频图标)
- uri(视频的url路径)
- group:(视频所属分组)
- type:(类型)



表:video_group(视频)

- type(类型)
- 

表:user(用户)

字段：

- id(唯一标识)
- username(用户名)
- password(密码)



表:admin(管理员)

字段:

- id(唯一标识)
- username(用户名)
- password(密码)
- rank(管理员级别)





表:comment(品论)

字段:

- 



