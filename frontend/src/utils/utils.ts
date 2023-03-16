export interface Example {
    input: string;
    output: string;
}

function stringToExamples(value: string) {
    let reg = "<input>([\\s\\S]*?)</input><output>([\\s\\S]*?)</output>";
    let re = RegExp(reg, "g");
    let objList:Array<Example> = []
    let tmp;
    while (tmp = re.exec(value)) {
        objList.push({ input: tmp[1], output: tmp[2] })
    }

    return objList
}

function examplesToString(objList: any) {
    if (objList.length == 0) {
        return "";
    }
    let result = ""
    for (let obj of objList) {
        result += "<input>" + obj.input + "</input><output>" + obj.output + "</output>"
    }
    return result
}

export default {
    stringToExamples: stringToExamples,
    examplesToString: examplesToString,
}