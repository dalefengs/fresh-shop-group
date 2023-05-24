import {parseDateStr} from '@/utils/date.js'

export const filters = {
    parseDate: value => {
        return parseDateStr(value);
    }
}
