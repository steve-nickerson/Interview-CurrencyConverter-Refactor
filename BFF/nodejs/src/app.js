import express from 'express';
import {CurrencyConverter} from './currencyConverter.js';
import * as currencyConverterRepository from './currencyConverterRepository.js';

const app = express();
const port = 58415;

app.get('/', (req, res) => {
    const example =
        `<html lang="en"><head><title>Sample Request</title></head>
            <body>
              <h2>Sample Request</h2>
              <a href="http://localhost:${port}/api/currencyConverter?from=USD&to=MXN&amount=100">http://localhost:${port}/api/currencyConverter?from=USD&to=MXN&amount=100</a>
            </body></html>`;
    res.send(example);
});

app.get('/api/currencyConverter', (req, res) => {
    const convertedAmount = CurrencyConverter(currencyConverterRepository).getConvertedAmount(req.query.from, req.query.to, req.query.amount);
    res.status(200).send(JSON.stringify(convertedAmount));
});

app.listen(port, (err) => {
    if (err) console.log(err);
    console.log(`Listening at http://localhost:${port}`);
});