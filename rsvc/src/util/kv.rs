use rocksdb::{Options, DB};

use crate::util;

pub fn get(key: String) -> String {
    // NB: db is automatically closed at end of lifetime
    let path = util::conf::Props::init().kv_dir;
    {
        let db = DB::open_default(path).unwrap();
        match db.get(key) {
            Ok(Some(value)) => {
                return String::from_utf8(value).unwrap();
            }
            Ok(None) => return String::from("{}"),
            Err(e) => {
                println!("{}",e);
                let err_out = r#"{"error":"error fetching records"}"#;
                return err_out.to_string();
            }
        }
    }
}

pub fn put(key: String, value: String) {
    // NB: db is automatically closed at end of lifetime
    let path = util::conf::Props::init().kv_dir;
    {
        let db = DB::open_default(path).unwrap();
        db.put(key, value).unwrap();
    }
}
