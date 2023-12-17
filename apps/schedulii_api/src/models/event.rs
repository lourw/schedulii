use serde::{Deserialize, Serialize};
use sqlx::types::chrono::{DateTime, Utc};
use sqlx::FromRow;

#[derive(Debug, FromRow, Serialize, Deserialize)]
pub struct Event {
    pub event_id: i32,
    pub event_name: String,
    pub start_time: DateTime<Utc>,
    pub end_time: DateTime<Utc>,
    pub location: String,
}

#[derive(Debug, FromRow, Serialize, Deserialize)]
pub struct NewEvent {
    pub event_name: String,
    pub start_time: DateTime<Utc>,
    pub end_time: DateTime<Utc>,
    pub location: String,
}
