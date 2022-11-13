use std::env;

pub struct Props {
    pub kafka_brokers: String,
    pub base_path: String,
    pub web_dir: String,
    pub web_address: String,
    pub kv_dir: String,
    pub test_default: String,
}

impl Props {
    pub fn init() -> Self {
        Props {
            kafka_brokers: set_broker(),
            base_path: set_base_path(),
            web_dir: set_web_dir(),
            web_address: set_web_address(),
            kv_dir: set_kv_dir(),
            test_default: Default::default(),
        }
    }
}

fn set_broker() -> String{
    return match env::var("KAFKA_BROKERS") {
            Ok(val) => val,
            Err(_e) => "172.18.0.41:9092".to_string(),
    };
}

fn set_base_path() -> String{
    return match env::var("BASE_PATH") {
            Ok(val) => val,
            Err(_e) => "/api/v1/".to_string(),
    };
}

fn set_web_dir() -> String{
    return match env::var("WEB_DIR") {
            Ok(val) => val,
            Err(_e) => "../ui/dist/".to_string(),
    };
}

fn set_web_address() -> String{
    return match env::var("WEB_ADDRESS") {
            Ok(val) => val,
            Err(_e) => "0.0.0.0:9000".to_string(),
    };
}

fn set_kv_dir() -> String{
    return match env::var("KV_DIR") {
            Ok(val) => val,
            Err(_e) => "/data/kv_def_dir".to_string(),
    };
}