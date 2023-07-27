using Microsoft.AspNetCore.Mvc;

namespace Interview_CurrencyConverter_Refactor.Controllers
{
    [Route("api/[controller]")]
    public class CurrencyConverterController : Controller
    {
        ICurrencyConverter converter;

        public CurrencyConverterController(ICurrencyConverter converter)
        {
            this.converter = converter;
        }

        [HttpGet]
        public decimal Get(string from, string to, decimal amount)
        {
            decimal convertedAmount = converter.GetConvertedAmount(from, to, amount);
            return convertedAmount;
        }
        
    }
}
