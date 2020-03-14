/* eslint-disable no-console */

function createChart(title, name, categories = [], data = [])
{
    const options = {
        chart: {
          type: 'line',
          background: '#424242',
          toolbar: {
              show: false
          },
          shadow: {
            enabled: false,
            color: 'red',
            top: 3,
            left: 2,
            blur: 3,
            opacity: 1
          },
        },
        theme: {
            mode: 'dark', 
            palette: 'palette1', 
            monochrome: {
                enabled: false,
                color: '#255aee',
                shadeTo: 'dark',
                shadeIntensity: 0.65
            },
        },
        xaxis: {
          type: 'string',
          categories: [],
          labels: {
              show: false
          }
        },
        fill: {
          type: 'gradient',
          gradient: {
            shade: 'dark',
            gradientToColors: ['#FDD835'],
            shadeIntensity: 1,
            type: 'horizontal',
            opacityFrom: 1,
            opacityTo: 1,
            stops: [0, 100, 100, 100]
          },
        }
    }

    title = ""
    options.xaxis.categories = categories

    return {
        datas: [{
            name: name,
            data: data
        }],
        options: options
    }
}

export async function priceChart()
{    
    const categories = []
    const data = []
    let response = await fetch("https://api.coingecko.com/api/v3/coins/dero/market_chart?vs_currency=usd&days=1").then(response => {
        return response.json()
    });
    let val = response.prices
    
    for (let i = 0; i < val.length; i++)
    {
        let m = new Date(val[i][0])
        var date = m.getUTCFullYear() +"/"+ (m.getUTCMonth()+1) +"/"+ m.getUTCDate() + " " + m.getUTCHours() + ":" + m.getUTCMinutes() + ":" + m.getUTCSeconds();
        categories.push(date)
        data.push(val[i][1].toFixed(5))
    }

    return createChart("Price 1 Day Chart", "USD", categories, data)
}