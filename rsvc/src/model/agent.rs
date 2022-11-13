use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct Agent {
    name: String,
    security_profile: String,
    id: String,
    e_tag: String,
    entity_id: String
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct CreateAgentRequest {
    name: String,
    security_profile: String,
    entity_id: String
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct UpdateAgentRequest {
    name: String,
    security_profile: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct OnlineStatus {
    status: String,
    since: String
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct UpdateDataSourceConfigurationRequest {
    configuration_id: String,
    data_sources: Vec<DataSource>
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct DataSourceConfiguration {
    id: String,
    e_tag: String,
    configuration_id: String,
    data_sources: Vec<DataSource>
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct DataSource {
    name: String,
    description: String,
    data_points: Vec<DataPoint>,
    custom_data: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct DataPoint {
    id: String,
    name: String,
    description: String,
    #[serde(rename = "type")]
    typed: String,
    unit: String,
    #[serde(rename = "camelCase")]
    custom_data: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct TokenKey {
    alg: String,
    #[serde(rename = "use")]
    used: String,
    value: String,
    e: String,
    n: String,
    kty: String,
    kid: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct AccessToken {
    access_token: String,
    token_type: String,
    expires_in: String,
    jti: String,
    scope: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Configuration {
    content: OnboardingConfigurationContent,
    expiration: String,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "camelCase")]
pub struct OnboardingConfigurationContent {
    base_url: String,
    iat: String,
    client_credential_profile: String,
    client_id: String,
    tenant: String
}

#[derive(Debug, Serialize, Deserialize)]
pub struct RotationKeys {
    client_id: String,
    jwks: Jwks
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Jwks {
    keys: Vec<Key>
}

#[derive(Debug, Serialize, Deserialize)]
pub struct Key {
    e: String,
    n: String,
    kty: String,
    kid: String,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ClientIdentifier {
    client_id: String,
    client_secret: String,
    client_secret_expires_at: String,
    grant_types: Vec<String>,
    registration_access_token: String,
    registration_client_uri: String,
    token_endpoint_auth_method: String
}

