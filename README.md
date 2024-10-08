Customer visits your site.
Customer chooses to make payment.
Your site creates a Bill via API call.
Billplz API returns Bill's URL.
Your site redirects the customer to Bill's URL.
The customer makes payment via payment option of choice.
Billplz sends a server-side update to your site upon payment failure or success. (Basic Callback URL / X Signature Callback URL depending on your configuration) [your backend server should capture the transaction update at this point] refer to X Signature Callback Url.
Billplz redirects (Payment Completion) the customer back to your site if redirect_url is not empty (Basic Redirect URL / X Signature Redirect URL depending on your configuration) [your server should capture the transaction update at this point and give your user an instant reflection on the page loaded] refer to X Signature Redirect Url or, The customer will see Billplz receipt if redirect_url is not present.

- [ ] Create a Bill via API call
- [ ] Billplz API returns Bill's URL
- [ ] Redirect the customer to Bill's URL
- [ ] Customer makes payment via payment option of choice
- [ ] Billplz sends a server-side update to your site upon payment failure or success
- [ ] Redirect the customer back to your site if redirect_url is not empty
- [ ] Customer will see Billplz receipt if redirect_url is not present

## Billplz API

- [ ] Create Bill
- [ ] Get Bill
- [ ] Get Bills
- [ ] Delete Bill
