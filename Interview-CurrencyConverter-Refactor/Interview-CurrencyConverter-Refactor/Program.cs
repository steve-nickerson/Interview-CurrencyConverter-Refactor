using Interview_CurrencyConverter_Refactor;

WebApplicationBuilder builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddCors(options => options
    .AddPolicy("all", p => p
        .WithOrigins("http://localhost:5173")
        .AllowAnyMethod()
        .WithExposedHeaders()
        .AllowAnyHeader()
        .Build()));

builder.Services.AddTransient<ICurrencyConverter, CurrencyConverter>();
builder.Services.AddTransient<ICurrencyConverterRepository, CurrencyConverterRepository>();

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

WebApplication app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseCors("all");

app.UseAuthorization();

app.MapControllers();

app.Run();