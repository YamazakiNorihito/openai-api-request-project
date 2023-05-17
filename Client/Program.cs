using System.Net.Http.Headers;
using Client;
using Client.Services;
using Microsoft.AspNetCore.Components.Web;
using Microsoft.AspNetCore.Components.WebAssembly.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

var builder = WebAssemblyHostBuilder.CreateDefault(args);
builder.RootComponents.Add<App>("#app");
builder.RootComponents.Add<HeadOutlet>("head::after");

builder.Services.AddHttpClient();
builder.Services.AddScoped(sp => new HttpClient { BaseAddress = new Uri(builder.HostEnvironment.BaseAddress) });


builder.Services.AddSingleton<IOpenAIService, OpenAIService>(
    p =>
    {
        var config = p.GetRequiredService<IConfiguration>();
        var httpClientFactory = p.GetRequiredService<IHttpClientFactory>();
        return new OpenAIService(httpClientFactory, config["BackEndUri"]!);
    }
);
await builder.Build().RunAsync();
