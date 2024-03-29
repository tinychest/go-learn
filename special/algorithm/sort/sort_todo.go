package sort

// http://www.360doc.com/content/13/0921/16/9615799_316009153.shtml
// https://lailin.xyz/post/51203.html

// 【概念】
//
// 排序算法固然多，种类丰富，但是选择上就和技术选型的概念是一样的：没有好和坏，只有最合适
//
// 排序算法的选择标准：数据量、数据特征（如数据大部分片段已经是有序的）
// 排序算法的算则，实际上这里边有个误区，就是算法是否好，是否快，只看算法的平均时间复杂度 - 其实也就是忽略了“平均”二字
// 从结论上说，如果数据量不大，且数据基本有序，插入排序吊打快速排序
//
// 【算法是否稳定】
// 假设 arr[i] = arr[j] 且 i < j，当排序完成后 arr[i] 所在的下标依旧小于 arr[j]

// 【分类】
// [插入排序]
// - 常见 直接插入、希尔排序
// - 不常见 Tree sort、Library sort、Patience sort
//
// [交换排序]
// - 常见 冒泡排序、快速排序
// - 不常见 鸡尾酒排序、奇偶排序、梳排序、Strand sort
//
// [选择排序]
// - 选择排序（Select Sort）
// - 堆排序（Heap Sort）
//
// [归并排序]
//  - 归并排序（Merge Sort）
//  - Strand sort
//
// <非比较排序>
//  - 计数排序（Counting Sort）
//  - 基数排序（Radix Sort）
//  - 桶排序（Bucket Sort） TODO
//
// <比较排序> TODO
//
// go sdk sort TODO
