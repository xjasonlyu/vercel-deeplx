# DeepLX Vercel

Host a serverless translation API on Vercel, powered by the official [DeepLX](https://github.com/OwO-Network/DeepLX) Go
implementation.

## Features

- **Serverless by Design**: Built on a serverless architecture for superior scalability and simplicity.
- **Effortless Deployment**: Deploy to Vercel with a single click and minimal configuration.
- **Local Development Support**: Fully supports local development for testing and customization.
- **Official DeepLX Compatibility**: Always stays in sync with the latest official `DeepLX` API.

## Deploying on Vercel

Deploy your DeepLX translation API by clicking the button below.

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fxjasonlyu%2Fvercel-deeplx)

> After clicking the button, simply follow the prompts to set up the API on Vercel.

## Running Locally

To run/test this API locally, you need to install [Vercel CLI](https://vercel.com/docs/cli) and Go, and then run the
following commands in your terminal:

```bash
git clone https://github.com/xjasonlyu/vercel-deeplx.git
cd vercel-deeplx
vercel dev
```

This will start a local dev server with the API running on `http://localhost:3000`.

## Environments

| Variable     | Description                           | Default |
|--------------|---------------------------------------|---------|
| `TOKEN`      | Access token to protect your API      | `null`  |
| `DL_SESSION` | DeepL Pro Account `dl_session` cookie | `null`  |
| `PROXY`	     | The http proxy server address         | 	`null` |

The above environment variables are optional and only need to be configured on Vercel when required.
See [variables](https://deeplx.owo.network/install/variables.html) for more details.

## Usage

Once deployed or running locally, you can start making requests to translate text.
For example, make a `POST` request to `/translate` with the following cURL command:

> Replace `https://your-deployment-url` with the actual URL of your Vercel deployment or `http://localhost:3000` if
> you're running locally.

```bash
curl -X POST https://your-deployment-url/translate \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your_access_token>" \
  -d '{
      "text": "Hello, world!",
      "source_lang": "EN",
      "target_lang": "DE"
  }'
```

For more detailed cURL usage and examples, see [here](https://deeplx.owo.network/integration/curl.html).

## Important Notes

- **Rate Limiting**: Excessive requests may result in temporary blocking by DeepL. Please avoid sending too many
  requests in a short period.
- **Usage Limitations**: This API is intended for personal or development use. For commercial use, consider subscribing
  to DeepL's official API.
- **No Guarantee**: As this API relies on DeepL's public endpoints, functionality may break if DeepL changes their API.
  Please report issues or contribute fixes if this occurs.

## Auto Update

- To manually update, you need to go to the **Actions** tab and run the workflow.
- This workflow is also automatically triggered on a daily basis.

## Credits

- [OwO-Network](https://github.com/OwO-Network) for the project [DeepLX](https://github.com/OwO-Network/DeepLX).
- [bropines](https://github.com/bropines) for inspiration [Deeplx-vercel](https://github.com/bropines/Deeplx-vercel).

## Contributing

Contributions are welcome! If you have a suggestion or fix, please fork the repository and submit a pull request.

## License

This project is open-sourced under the MIT license. See the [LICENSE](LICENSE) file for more details.
