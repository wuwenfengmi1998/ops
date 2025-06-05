# l-calendar

# ## 使用方法
配置easycom规则后，自动按需引入，无需`import`组件，直接引用即可。

```html
<template>
	<l-calendar v-model="show" @change="change" :isRange="true"></l-calendar>
</template>
```
## 组件属性

| 属性 | 类型 | 默认值 | 说明 |
|:---:|:---:|:---:|---|
| value | Boolean | false | v-module 双向绑定的值 |
| isRange | Boolean | false | 是否选择范围 true是  false选择单个日期 |
| maxYear | Number | 2100 | 可切换最大年份 |
| minYear | Number | 1920 | 可切换最小年份 |
| minDate | String | '1920-01-01' | 最小可选日期 不在范围内日期禁选 |
| maxDate | String | '2100-1-1' | 最大可选日期 不在范围内日期禁选 |
| title | String | '日期选择' | 组件标题 |
| monthChangeColor | String | #999 | 月份切换箭头颜色 |
| yearChangeColor | String | '#bfbfbf' | 年份切换箭头颜色 |
| color | String | #333 | 默认日期字体颜色 |
| activeColor | String | '#fff' | 选中日期字体颜色 |
| activeBgColor | String | '#55BBF9' | 选中日期背景色 |
| rangeBgColor | String | rgba(85, 187, 249, 0.1) | 范围内日期背景色 |
| startText | String | 开始 | 范围选择时生效 开始日期自定义文字 |
| endText | String | 结束 | 范围选择时生效 结束日期自定义文字 |
| lunar | Boolean | true | 是否显示农历 |
| initStartDate | String | '' | 初始化开始选中日期 格式： 2020-06-06 或 2020/06/06 |
| initEndDate | String | '' | 初始化开始选中日期 格式： 2020-06-06 或 2020/06/06 |


## 组件事件

| 名称 | 触发时机 |
|:---:|---|
| change | 点击确定按钮选择日期 |

