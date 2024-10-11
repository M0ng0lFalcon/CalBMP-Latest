/*
    date1 >  date2 : return true
    date1 <= date2 : return false
*/
export function CompareDate(date1_str, date2_str) {
  const date1 = new Date(date1_str);
  const date2 = new Date(date2_str);

  // console.log("data1:", date1.toString());
  // console.log("date2:", date2.toString());

  if (date1 > date2) {
    return true;
  } else if (date1 <= date2) {
    return false;
  }
}

export function GetIdxOfDateList(DateList, date_str) {
  date_str = date_str.replace(/-/, "/");
  const len = DateList.length;
  const curDate = new Date(date_str);
  let i = 0;
  for (i; i < len; i += 1) {
    const tmpDate = new Date(DateList[i]);
    if (tmpDate > curDate) {
      return i;
    }
  }
  return i;
}
