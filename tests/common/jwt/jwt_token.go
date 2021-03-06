// Copyright 2019 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package jwt includes sample JWT Token used in e2e tests.
package jwt

const (
	// Payload {
	//  "exp": 4715782722,
	//  "groups": [
	//    "group-1"
	//  ],
	//  "iat": 1562182722,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-1@istio.io --sub=sub-1 --listclaim groups "group-1"
	// nolint: lll
	TokenIssuer1 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3MTU3ODI3MjIsImdyb3VwcyI6WyJncm91cC0xIl0sImlhdCI6MTU2MjE4MjcyMiwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.JUWnHmKP01ZJMxtE8ke7Lv7G_eXL9yEoa4xWi__Sk09IQsuBGnhK75nQfl4aJ1fLthu-AvIp-GZe-1hDnBHhO4MQ7tXYeCdgnFkVGvhvZwOWpOl79H2U-mRX1eyjsHG0ZzwHwWMMLpPW6HQYCbTCEkZKRtoFoSum0-ilGqLO7EGRMEFpKYKnZV0N9pZIgX8uPlHoOVCZjqvhDtFjP-PG89vvNozeV7u47LMMzQeQtzCDmbRrIw1KVHZ885N1qiLcf-RNR9x9UINK_d98QKqkYsJtI2QTdVzevzSF8LfvaGLh75tvAL8h7L221haAKEnICuRqGWNO8qIwsj0k89kdGA"

	// Payload {
	//  "exp": 4715782783,
	//  "groups": [
	//    "group-2"
	//  ],
	//  "iat": 1562182783,
	//  "iss": "test-issuer-2@istio.io",
	//  "sub": "sub-2"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=3153600000 --iss=test-issuer-2@istio.io --sub=sub-2 --listclaim groups "group-2"
	// nolint: lll
	TokenIssuer2 = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjQ3MTU3ODI3ODMsImdyb3VwcyI6WyJncm91cC0yIl0sImlhdCI6MTU2MjE4Mjc4MywiaXNzIjoidGVzdC1pc3N1ZXItMkBpc3Rpby5pbyIsInN1YiI6InN1Yi0yIn0.KmOIdRoI-FfaWR6xdMWrguOypmqD5UC2VR6B1Cw8V0aBhaqpCBcn5fDvA13mSiU3sx3VFF7-CbnnZVz4j6k2RsYNnWDDJTzr4GnxIXJYJpIRBcE0Au9mcZM0cVhW9BBtKtc0twy0I5S7H7qIAEPBq6iYm8IED409K7YIuI92kLeSe--ehhs1yWmI-YbixYwwq7viD1ZIRdbP5b2ZxU6A0VHnWVXFBGSoYshSLdJS6fnlU480SXyz0oUakDpkaVkz6n9LXolggs6Cz6Gq9OIEs2zUz2DqAJ5h6HChjIQ8AydwQnRMx4qnEbdtqq2qy1h5HHfosq8pcGbiQKf7CSQKHw"

	// Payload {
	//  "exp": 1562182856,
	//  "groups": [
	//    "group-1"
	//  ],
	//  "iat": 1562182855,
	//  "iss": "test-issuer-1@istio.io",
	//  "sub": "sub-1"
	// }
	// Generated by: security/tools/jwt/samples/gen-jwt.py tests/common/jwt/key.pem -jwks=tests/common/jwt/jwks.json
	// --expire=1 --iss=test-issuer-1@istio.io --sub=sub-1 --listclaim groups "group-1"
	// nolint: lll
	TokenExpired = "eyJhbGciOiJSUzI1NiIsImtpZCI6InRUX3c5TFJOclk3d0phbEdzVFlTdDdydXRaaTg2R3Z5YzBFS1I0Q2FRQXciLCJ0eXAiOiJKV1QifQ.eyJleHAiOjE1NjIxODI4NTYsImdyb3VwcyI6WyJncm91cC0xIl0sImlhdCI6MTU2MjE4Mjg1NSwiaXNzIjoidGVzdC1pc3N1ZXItMUBpc3Rpby5pbyIsInN1YiI6InN1Yi0xIn0.o9hp0P5DS66Q7wP38wGo0AKS5HoWdHXrAUdNLnXzVC4MwU4N9o3U0PfDgWp8A7YIbDuuQAtKJyCHEJCV3JEh7Xb5Tz_12hYQcXcJ0FTA6pOXrbWRjkAVMhs3-OHiKt855s39l2OKrLwmJ3ph7LV4z8J8i-2LE2hQH18R00_dcx2agoY1VNYH2LSErBYx6e-HQlKFW8c9sQ1CHG7u4ns1I2e23A0nBlRWRUHUYIMo6sfAPWhyQWl1kpRzg6b3fyXGfUpgeEmdIPDK7MfRUZA-0epFGjvoqCwgMdEEQ-O_pH5y2qV1jPpu-9IO_FdpYhHofKMTn_ql05ys6zoIHj9Gng"

	TokenInvalid = "TokenInvalid"
)
