const CurrencySelection = ({ currencies, selectedValue, onCurrencySelected }) => {
    return (
        <select
            value={selectedValue}
            onChange={e => onCurrencySelected(e.target.value)}
        >
            {currencies.map(c => (
                <option key={c.code} value={c.code}>
                    {c.name} ({c.code})
                </option>
            ))}
        </select>
    );
};

export default CurrencySelection;