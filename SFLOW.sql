-- table of flow items , and timestamp
create table SFLOW (
 ts TIMESTAMP DEFAULT NOW not null,
 agent varchar(64) not null,
 inport integer not null,
 outport integer not null,
 srcip varchar(64) not null,
 dstip varchar(64) not null,
 ipprotocol integer not null,
 srcport integer not null,
 dstport integer not null,
 tcp_flags integer not null,
 packet_size integer not null,
 ip_size integer not null,
 sampling_rate integer not null) USING TTL 10 MINUTES ON COLUMN ts;
-- need to use TTL
create index idx_sflow_ts on SFLOW (ts);
