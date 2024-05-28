import sinon from 'sinon';
import {expect} from 'chai';

import {currencyConversion} from "../src/currencyConverterRepository.js";
import {CurrencyConverter} from "../src/currencyConverter.js";

describe('CurrencyConverter', function () {
    let currencyConversions;
    let currencyConverter;
    let mockRepo;

    beforeEach(function () {
        currencyConversions = [
            currencyConversion('CAD', 'Canada Dollars', 2),
            currencyConversion('MXN', 'Mexico Pesos', 20)
        ];

        mockRepo = {
            getConversions: sinon.stub().returns(currencyConversions)
        };

        currencyConverter = CurrencyConverter(mockRepo);
    });

    it('GetConvertedAmount_ConvertsCorrectly', function () {
        const actual = currencyConverter.getConvertedAmount('CAD', 'MXN', 10);
        expect(actual).to.equal(100);
    });

    it('GetConvertedAmount_ConvertsCorrectlyWithMismatchedCase', function () {
        const actual = currencyConverter.getConvertedAmount('caD', 'mXn', 10);
        expect(actual).to.equal(100);
    });

    it('GetConvertedAmount_UnknownFromCountryCodeThrowsException', function () {
        expect(() => currencyConverter.getConvertedAmount('XXX', 'MXN', 10)).to.throw();
    });

    it('GetConvertedAmount_UnknownToCountryCodeThrowsException', function () {
        expect(() => currencyConverter.getConvertedAmount('CAD', 'XXX', 10)).to.throw();
    });
});