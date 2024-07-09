# 多源连接

目前，连接是 eKuiper 中合并多个数据源的唯一方法。它需要一种方法来对齐多个来源并触发连接结果。

eKuiper支持的连接包括：

- 多流的连接：必须在一个窗口中进行。
- 流和表的连接：流将是连接操作的触发器。

eKuiper 支持的连接类型包括 LEFT、RIGHT、FULL 和 CROSS 。