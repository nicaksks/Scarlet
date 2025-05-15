use log::info;
use models::response::BsResponse;
use reqwest::{Method, header::HeaderMap};
use serde::de::DeserializeOwned;
use std::fmt::Debug;

use crate::models;
pub struct Client;

const BASE_URL: &'static str = "http://localhost:3000/v1";

impl Client {
    async fn instance(&self) -> reqwest::Client {
        let client = reqwest::Client::builder();
        client.default_headers(self.headers()).build().unwrap()
    }

    fn headers(&self) -> HeaderMap {
        let mut header = HeaderMap::new();
        header.insert("Content-Type", "application/json".parse().unwrap());
        header.insert("X-Scarlet-Session", "".parse().unwrap());
        header
    }

    async fn client<T>(&self, method: Method, endpoint: String) -> Result<T, String>
    where
        T: DeserializeOwned,
        T: Debug,
    {
        let client: reqwest::Client = self.instance().await;

        let response = client
            .request(method, format!("{}{}", BASE_URL, endpoint))
            .send()
            .await
            .unwrap();

        let data = response.json::<BsResponse<T>>().await;
        println!("{:?}", data);

        if let Ok(data) = data {
            if data.code != 200 || data.result.is_none() {
                info!("[Scarlet HTTP] ERROR -> {:?}", data.message);
                return Err(data.message);
            }

            return Ok(data.result.unwrap());
        }

        info!("[Scarlet HTTP] ERROR -> {:?}", data);
        Err("unknown".to_string())
    }

    pub async fn get<T>(&self, endpoint: String) -> Result<T, String>
    where
        T: DeserializeOwned,
        T: Debug,
    {
        self.client(Method::GET, endpoint).await
    }
}
