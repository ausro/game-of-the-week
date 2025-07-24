export default function convDateToNum(date: string) {
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
    let args: string[] = date.split(" ");

    const month = months.indexOf(args[0]).toString();
    const year = args[2];
    const day = args[1];

    const fullDateStr = year + month + day;
    return Number.parseInt(fullDateStr);
}
