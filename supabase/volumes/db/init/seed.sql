CREATE USER readonly WITH ENCRYPTED PASSWORD 'n49dK5n7MW93OG';
GRANT pg_read_all_data TO readonly;

INSERT INTO storage.buckets (id, name) VALUES ('test', 'test');
CREATE POLICY "Select 2487m_0" ON storage.objects FOR SELECT TO public USING (bucket_id = 'test');
CREATE POLICY "Insert 2487m_1" ON storage.objects FOR INSERT TO public WITH CHECK (bucket_id = 'test');