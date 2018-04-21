/**
 * Created by zhoulin on 2017/6/7.
 */
function isPositiveInteger(s){//是否为正整数
    var reg = /^[1-9]\d*$/;
    return reg.test(s)
}