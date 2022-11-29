import { useEffect } from 'react';
import * as echarts from 'echarts';

export default function Skills(props) {

    useEffect(() => {
        drawSkills(props.theme)
    });


    return (
        <div className="pure-g" id="skills">
            {Array.from(skills.keys()).map((key) =>
                <div className="pure-u-1 pure-u-md-1-3 skillsItem" id={key} key={key}></div>
            )}
        </div>
    );
}

const skills = new Map([
    ['MicroService', 70],
    ['Observability', 75],
    ['Streaming', 75],
    ['NLP', 60],
    ['RxJava', 60],
    ['OpenCV', 60]
]);


function drawSkills(theme) {
    for (const [key, value] of skills.entries()) {
        drawSkill(theme, key, key, value)
    }
}

function drawSkill(theme, elementId, chartName, score) {
    var myChart;
    var chartDom = document.getElementById(elementId);
    echarts.dispose(chartDom);
    if (theme === 'dark') {
        myChart = echarts.init(chartDom, 'dark');
    } else {
        myChart = echarts.init(chartDom);
    }
    var option;

    option = {
        tooltip: {
            formatter: '{b} : {c}%'
        },
        series: [
            {
                name: chartName,
                type: 'gauge',
                progress: {
                    show: true
                },
                axisLabel: {
                    fontSize: 10
                },
                title: {
                    show: true,
                    fontSize: 11
                },
                detail: {
                    valueAnimation: true,
                    formatter: '{value}',
                    fontSize: 1
                },
                data: [
                    {
                        value: score,
                        name: chartName,
                        fontSize: 1
                    }
                ]
            }
        ]
    };

    option && myChart.setOption(option);
}
