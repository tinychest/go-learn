package _status

/*
【状态模式】
将事件触发的状态转移和动作执行，拆分到不同的状态类中，避免分支判断逻辑

【有限状态机 FSM 实现方式】
分支逻辑：状态和逻辑较少时可以使用
查表法：状态很多、状态转移比较复杂的状态机，查表发比较合适
状态模式：适用于状态不多、状态转移也比较简单；但是事件触发执行的动作包含的业务逻辑可能比较复杂的状态机
*/
