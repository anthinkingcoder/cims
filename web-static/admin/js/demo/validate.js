/**
 * Created by zhoulin on 2017/6/7.
 */
function isPositiveInteger(s){//是否为正整数
    var re = /^[0-9]+$/ ;
    return re.test(s)
}