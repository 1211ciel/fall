> 在im目录下执行 goctl rpc proto -I . -src comet/comet.proto -dir comet/ -style goZero

## ciel-im

整体思想参考goim. 实现方式使用 go-zero redis

comet 负责奖励和维持客服端的长连接

logic 提供三种维度的消息(全局,room,用户)投递,还包括业务逻辑,Session管理

job 通过redis订阅发布功能 进行要洗的分发

## 用户连接

1. 用户 发起ws请求到comet层
2. comet层 rpc到logic层
3. logic层 处理并返回结果
4. comet层 保存用户状态并返回结果

## 消息流转

1. 客服端 http请求发送消息到 logic层
2. logic层 处理处理完之后 异步提交到队列(存储,消峰)
3. job层 每个Job成员都从队列中消费消息,再投递给一个或者多个comet
4. comet层 成员收到消息后 再发送给客户端

用户channel

```go

```

消息结构体

```protobuf
message pushMsg{
  enum Type{
    PUSH = 0; // 单用户推送
    ROOM = 1; // 房间推送
    BROADCAST = 2; // 全频道广播
  }
  Type type = 1;
  int32 operation = 2;
  string server = 3;
  string room = 4; // 房间id, 后面广播时,发现用户如果监听了该房间号则推送消息
  bytes msg = 6;
}
```

## 生成消息

消息(初鉴权/心跳等基础数据包外)生成都是由logic完成第一手处理; logic提供HTTP接口以支持消息发送能力,主要由三个维度:用户,房间,全应用广播.