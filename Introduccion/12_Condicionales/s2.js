const formatter = (o) => {
    const res = o.reduce((acc, { id, question_name, question_value: value }) => {
        if (acc[question_name]) {
            acc[question_name].push({ id, value })
        } else {
            acc[question_name] = [{ id, value }]
        }
        return acc
    }, {})

    console.log(res);
}

const exampleJSON = [
    { 'id': 1, 'question_name': "What is your name?", 'question_value': "Jack" },
    { 'id': 2, 'question_name': "What is your hobby?", 'question_value': "Rugby" },
    { 'id': 3, 'question_name': "What is your name?", 'question_value': "Peter" },
    { 'id': 4, 'question_name': "What is your hobby?", 'question_value': "Tennis" },
    { 'id': 5, 'question_name': "What is your hobby?", 'question_value': "Basquetball" },
    { 'id': 6, 'question_name': "What is your hobby?", 'question_value': "Tennis" },
    { 'id': 7, 'question_name': "What is your hobby?", 'question_value': "Ping Pong" },
    { 'id': 8, 'question_name': "What is your hobby?", 'question_value': "Football" }
];

console.log(formatter(exampleJSON));