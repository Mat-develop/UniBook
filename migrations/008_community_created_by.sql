ALTER TABLE community
  ADD COLUMN created_by INT NULL,
  ADD CONSTRAINT fk_community_creator
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL;
