using System.Text;
using System.Text.Json;

namespace Client.Services;

public interface IOpenAIService
{
    Task<string> GetCompletionAsync(string prompt);
}

public class OpenAIService : IOpenAIService
{
    private readonly IHttpClientFactory _httpClientFactory;
    private readonly string _backEndUri;

    public OpenAIService(IHttpClientFactory httpClientFactory, string backEndUri)
    {
        (_httpClientFactory, _backEndUri) = (httpClientFactory, backEndUri);
    }

    public async Task<string> GetCompletionAsync(string prompt)
    {
        var requestData = new
        {
            prompt
        };
        var jsonRequestData = JsonSerializer.Serialize(requestData);
        var httpRequestMessage = new HttpRequestMessage(
            HttpMethod.Post,
            $"{_backEndUri}/completions")
        {
            Content = new StringContent(jsonRequestData, Encoding.UTF8, "application/json")
        };

        var httpClient = _httpClientFactory.CreateClient("OpenAIApi");
        using var response = await httpClient.SendAsync(httpRequestMessage, HttpCompletionOption.ResponseContentRead);
        var responseContent = await response.Content.ReadAsStringAsync();
        var responseObject = JsonSerializer.Deserialize<TextCompletionResponse>(responseContent,
            new JsonSerializerOptions
            {
                PropertyNameCaseInsensitive = true
            });

        return responseObject!.Choices[0].Text;
    }
}

public class TextCompletionResponse
{
    public string Id { get; set; }
    public string Object { get; set; }
    public long Created { get; set; }
    public string Model { get; set; }
    public List<Choice> Choices { get; set; }
    public Usage Usage { get; set; }
}

public class Choice
{
    public string Text { get; set; }
    public int Index { get; set; }
    public object Logprobs { get; set; }
    public string FinishReason { get; set; }
}

public class Usage
{
    public int PromptTokens { get; set; }
    public int CompletionTokens { get; set; }
    public int TotalTokens { get; set; }
}