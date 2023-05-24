/**
 * 格式化时间字符串
 * @param dateStr
 * @returns {*}
 */
export function parseDateStr(dateStr) {
    const date = new Date(dateStr);
    return formatDateTime(date, 'yyyy-MM-dd HH:mm:ss')
}


/**
 * 通用格式化时间
 * @param date
 * @param format
 * @returns {*}
 */
function formatDateTime(date, format) {
    const o = {
        'M+': date.getMonth() + 1, // 月份
        'd+': date.getDate(), // 日
        'h+': date.getHours() % 12 === 0 ? 12 : date.getHours() % 12, // 小时
        'H+': date.getHours(), // 小时
        'm+': date.getMinutes(), // 分
        's+': date.getSeconds(), // 秒
        'q+': Math.floor((date.getMonth() + 3) / 3), // 季度
        S: date.getMilliseconds(), // 毫秒
        a: date.getHours() < 12 ? '上午' : '下午', // 上午/下午
        A: date.getHours() < 12 ? 'AM' : 'PM', // AM/PM
    };
    if (/(y+)/.test(format)) {
        format = format.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length));
    }
    for (let k in o) {
        if (new RegExp('(' + k + ')').test(format)) {
            format = format.replace(
                RegExp.$1,
                RegExp.$1.length === 1 ? o[k] : ('00' + o[k]).substr(('' + o[k]).length)
            );
        }
    }
    return format;
}
