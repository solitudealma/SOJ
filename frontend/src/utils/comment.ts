import { Comment } from "@/api/types/comment";

/**
 * @description 获取各个层级的评论数据，保存在map中
 * @param {Array} allComments 所有评论数据
 * @return {Map<number, Comment[]>} levelDataMap 处理后的结果map
 */
function getEveryLevelData(allComments: Array<Comment>) {
  const levelDataMap = new Map<number, Array<Comment>>([
    [1, []],
    [2, []]
  ]);
  allComments.map(item => {
    switch (item.layer) {
      case 1:
        levelDataMap.get(1)!.push(item);
        break;
      case 2:
        levelDataMap.get(2)!.push(item);
        break;
    }
  });
  return levelDataMap;
}

/**
 * @description 将数据son转换为map，key是数组father的comment_id
 * @param {Array} parent 待转换的数组
 * @param {Array} son 待转换的数组
 * @return {Map<any, any>} tempMap 返回的map
 */
function convertToMap(parent: Array<Comment>, son: Array<Comment>) {
  // 获取所有key
  const keys = parent.map(item => item.commentId);
  // 声明map
  let tempMap = new Map<number, Array<Comment>>();
  // 初始化map
  for (const key of keys) {
    tempMap.set(key, []);
  }
  // 往map里塞数据
  son.map((item: Comment) => {
    // 因为只有两层，所以会出现子评论回复的不是根评论，而已某一个子评论，所以只要根评论相同就放入同一个根评论的回复中
    if (keys.includes(item.rootCommentId)) {
      tempMap.get(item.rootCommentId)!.push(item);
    }
  });
  return tempMap;
}


/**
 * @description 将所有子集评论塞到父级评论下，一共有五层数据
 * @param {Map<any, any>} everyLevelCommentsMap
 * @return {Array} resultArr 转换为树形结构的评论列表
 */
function convertCommentListToTree(everyLevelCommentsMap: Map<number, Array<Comment>>) {
  const resultArr: Array<Comment> = [];
  const level2Map = convertToMap(everyLevelCommentsMap.get(1)!, everyLevelCommentsMap.get(2)!);
  // 1层
  everyLevelCommentsMap.get(1)!.map(item => {
    let leve1Obj: Comment = {
      ...item,
    }
    if (level2Map.has(item.commentId)) {
      // 2层
      leve1Obj.reply = level2Map.get(item.commentId)!;
    }
    resultArr.push(leve1Obj);
  });
  return resultArr;
}

/**
 * @description 扁平化树形评论，为了分页展示
 * @param {Array<Comment>} commentArr 树形评论，共有5层
 * @param {Array<Comment>} resultArr 递归传入的当前结果值，最终遍历结束返回出去
 * @return {Array<Comment>} resultArr 扁平化后的评论列表
 */
function flatCommentArr(commentArr: Array<Comment>, resultArr = new Array<Comment>()) {
  commentArr.map(item => {
    let tempObj = <Comment>{};
    // push除了reply的所有字段
    Object.keys(item).map((key) => {
      if (key !== 'reply') {
        tempObj[key] = item[key];
      }
    });
    resultArr.push(tempObj);
    // 递归push数据，扁平化
    if (item.reply.length > 0) {
      flatCommentArr(item.reply, resultArr);
    }
  });
  return resultArr;
}


/**
 * @description 对评论数据list重新排序
 * @param {Array<Comment>} list 数据库请求过来的所有评论列表，已经按照时间降序排好序了
 * @return {Array<Comment>} 重新排序的评论列表，可以达到gitee的效果
 *
 * 仿写的gitee的评论列表，每次分页加载的数据条数包括子评论，
 * 那么我就把数据库的所有评论先转换为对应的树形结构，
 * 然后再扁平化，这样每次分页取数据就直接从扁平化后的列表中取即可，
 * 然后将取除的列表再转换为树形结构，就达到了gitee的效果
 *
 * 1、先找到各层级的评论
 *      1.1、    1层数据 commentLayer = 1
 *      1.2、    2层数据 commentLayer = 2
 *      1.3、    3层数据 commentLayer = 3
 *      1.4、    4层数据 commentLayer = 4
 *      1.5、    5层数据 commentLayer = 5
 * 2、将子集评论塞到父级评论下
 * 3、再将树形结构扁平化，变成列表结构，这样就可以根据这个列表进行分页了
 */

export const groupCommentList = function (list: Array<Comment>) {
  // 1、找到各层级的评论
  const everyLevelCommentsMap = getEveryLevelData(list);
  // 2、将子集评论递归塞到父级评论下
  const commentTree = convertCommentListToTree(everyLevelCommentsMap);
  // 3、扁平化数组 变成一维
  return flatCommentArr(commentTree);
}


/**
 * @description 将扁平化后的数据转换为树形结构
 * @param {Array<Comment>} list 扁平化完的评论列表，已经满足要求了
 * @return {Array<Comment>} 最终渲染到前台的树形结构评论
 *
 */
export const flatCommentListToTree = (list: Array<Comment>) => {
  // 1、找到各层级的评论
  const everyLevelCommentsMap = getEveryLevelData(list);
  // 2、将子集评论递归塞到父级评论下
  return convertCommentListToTree(everyLevelCommentsMap);
}

/*================= 评论数据分页处理 end =====================*/
