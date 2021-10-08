import React, { Component, Fragment } from 'react';

import Chart from "react-apexcharts";

class Grafica extends Component {

    constructor(props) {
        super(props);
        this.state = {

            series: [
                {
                    data: this.props.data
                }
            ],
            options: {
                colors: ['#ffffff'],
                chart: {
                    id: 'realtime',
                    height: this.props.height,
                    type: 'line',
                    animations: {
                        enabled: true,
                        easing: 'linear',
                        dynamicAnimation: {
                            speed: 1000
                        }
                    },
                    toolbar: {
                        show: true
                    },
                    zoom: {
                        enabled: false
                    }
                },
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    curve: 'smooth'
                },
                markers: {
                    size: 0
                },
                xaxis: {
                    categories: []
                },
                yaxis: {
                    max: this.props.max
                },
                legend: {
                    show: false
                },
            }
        };

    }

    componentDidMount() {

        this.setState({
            series: [
                {
                    data: this.props.data
                }
            ],
            options: {
                colors: ['#ffffff'],
                chart: {
                    id: 'realtime',
                    height: this.props.height,
                    type: 'line',
                    animations: {
                        enabled: true,
                        easing: 'linear',
                        dynamicAnimation: {
                            speed: 1000
                        }
                    },
                    toolbar: {
                        show: true
                    },
                    zoom: {
                        enabled: true
                    }
                },
                dataLabels: {
                    enabled: false
                },
                stroke: {
                    curve: 'smooth'
                },
                markers: {
                    size: 0
                },
                xaxis: {
                    labels: {
                        style: {
                            colors: "#ffffff"
                        }
                    }
                },
                yaxis: {
                    max: this.props.max,
                    labels: {
                        style: {
                            colors: ['#ffffff']
                        }
                    }
                },
                legend: {
                    show: false
                },
                tooltip: {
                    enabled: true,
                    theme: "dark"
                }
            }
        })
    }

    render() {
        return (

            <Fragment>
                <Chart options={this.state.options} series={this.state.series} type="line" width={this.props.width} />
            </Fragment>

        );
    }
}


export default Grafica