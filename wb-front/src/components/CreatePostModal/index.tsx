import React, { useCallback, useRef, useState } from 'react';
import { Button, Form, Input, Modal, Select, Space, Tag, Tooltip } from 'antd';
import { BoldOutlined, ItalicOutlined, LinkOutlined, PlusOutlined } from '@ant-design/icons';
import type { TextAreaRef } from 'antd/es/input/TextArea';
import { toast } from 'react-toastify';
import { createPost, getTags, type Tag as ApiTag } from '../../utils/api';
import styles from './createPostModal.module.scss';

interface Props {
  communityId: number;
  open: boolean;
  onClose: () => void;
  onSuccess: () => void;
}

const CreatePostModal: React.FC<Props> = ({ communityId, open, onClose, onSuccess }) => {
  const [form] = Form.useForm();
  const [body, setBody] = useState('');
  const [bodyError, setBodyError] = useState(false);
  const [links, setLinks] = useState<string[]>([]);
  const [linkInput, setLinkInput] = useState('');
  const [tagOptions, setTagOptions] = useState<ApiTag[]>([]);
  const [fetching, setFetching] = useState(false);
  const [submitting, setSubmitting] = useState(false);
  const bodyRef = useRef<TextAreaRef>(null);
  const debounceRef = useRef<ReturnType<typeof setTimeout> | null>(null);

  const getNativeTextarea = () =>
    bodyRef.current?.resizableTextArea?.textArea ?? null;

  const wrapSelection = (before: string, after = before) => {
    const el = getNativeTextarea();
    if (!el) return;
    const s = el.selectionStart;
    const e = el.selectionEnd;
    const newVal = body.slice(0, s) + before + body.slice(s, e) + after + body.slice(e);
    setBody(newVal);
    setTimeout(() => {
      el.selectionStart = s + before.length;
      el.selectionEnd = e + before.length + (e - s);
      el.focus();
    });
  };

  const prependLine = (prefix: string) => {
    const el = getNativeTextarea();
    if (!el) return;
    const s = el.selectionStart;
    const lineStart = body.lastIndexOf('\n', s - 1) + 1;
    const hasPrefix = body.slice(lineStart).startsWith(prefix);
    let newVal: string;
    let offset: number;
    if (hasPrefix) {
      newVal = body.slice(0, lineStart) + body.slice(lineStart + prefix.length);
      offset = -prefix.length;
    } else {
      newVal = body.slice(0, lineStart) + prefix + body.slice(lineStart);
      offset = prefix.length;
    }
    setBody(newVal);
    setTimeout(() => {
      const pos = Math.max(0, s + offset);
      el.selectionStart = pos;
      el.selectionEnd = pos;
      el.focus();
    });
  };

  const handleAddLink = () => {
    const url = linkInput.trim();
    if (!url) return;
    try { new URL(url); } catch {
      toast.error('Enter a valid URL (e.g. https://...)');
      return;
    }
    if (!links.includes(url)) setLinks((prev) => [...prev, url]);
    setLinkInput('');
  };

  const handleTagSearch = useCallback((value: string) => {
    if (debounceRef.current) clearTimeout(debounceRef.current);
    debounceRef.current = setTimeout(async () => {
      setFetching(true);
      try { setTagOptions(await getTags(value)); }
      finally { setFetching(false); }
    }, 300);
  }, []);

  const reset = () => {
    form.resetFields();
    setBody('');
    setLinks([]);
    setLinkInput('');
    setBodyError(false);
  };

  const handleOpen = () => {
    reset();
    getTags().then(setTagOptions).catch(() => {});
  };

  const handleSubmit = async () => {
    let values: { title: string; tags?: string[] };
    try { values = await form.validateFields(); } catch { return; }

    if (!body.trim()) { setBodyError(true); return; }
    setBodyError(false);
    setSubmitting(true);
    try {
      await createPost({
        communityId,
        title: values.title.trim(),
        body: body.trim(),
        links,
        tags: values.tags ?? [],
      });
      toast.success('Post created!');
      reset();
      onSuccess();
    } catch {
      toast.error('Failed to create post');
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Modal
      title="Create Post"
      open={open}
      onOk={handleSubmit}
      onCancel={() => { reset(); onClose(); }}
      okText="Post"
      confirmLoading={submitting}
      afterOpenChange={(visible) => visible && handleOpen()}
      width={660}
      destroyOnClose
    >
      <Form form={form} layout="vertical">
        <Form.Item name="title" label="Title" rules={[{ required: true, message: 'Title is required' }]}>
          <Input placeholder="Post title" maxLength={200} />
        </Form.Item>

        <Form.Item
          label="Body"
          required
          validateStatus={bodyError ? 'error' : ''}
          help={bodyError ? 'Body is required' : ''}
        >
          <div className={styles.editorWrapper}>
            <div className={styles.toolbar}>
              <Tooltip title="Bold"><Button size="small" icon={<BoldOutlined />} onClick={() => wrapSelection('**')} /></Tooltip>
              <Tooltip title="Italic"><Button size="small" icon={<ItalicOutlined />} onClick={() => wrapSelection('*')} /></Tooltip>
              <div className={styles.separator} />
              <Tooltip title="Heading 1"><Button size="small" onClick={() => prependLine('# ')}>H1</Button></Tooltip>
              <Tooltip title="Heading 2"><Button size="small" onClick={() => prependLine('## ')}>H2</Button></Tooltip>
              <Tooltip title="Heading 3"><Button size="small" onClick={() => prependLine('### ')}>H3</Button></Tooltip>
            </div>
            <Input.TextArea
              ref={bodyRef}
              variant="borderless"
              value={body}
              onChange={(e) => { setBody(e.target.value); setBodyError(false); }}
              placeholder={"Write your post…\n\nSupports **bold**, *italic*, # Heading"}
              rows={7}
              maxLength={10000}
              showCount
            />
          </div>
        </Form.Item>

        <Form.Item label="Links">
          <Space.Compact style={{ width: '100%' }}>
            <Input
              value={linkInput}
              onChange={(e) => setLinkInput(e.target.value)}
              onPressEnter={handleAddLink}
              placeholder="https://example.com"
              prefix={<LinkOutlined />}
            />
            <Button icon={<PlusOutlined />} onClick={handleAddLink}>Add</Button>
          </Space.Compact>
          {links.length > 0 && (
            <div className={styles.linkChips}>
              {links.map((url) => {
                let host = url;
                try { host = new URL(url).hostname; } catch {}
                return (
                  <Tag
                    key={url}
                    closable
                    onClose={() => setLinks((prev) => prev.filter((l) => l !== url))}
                    icon={<LinkOutlined />}
                    color="geekblue"
                  >
                    {host}
                  </Tag>
                );
              })}
            </div>
          )}
        </Form.Item>

        <Form.Item name="tags" label="Tags">
          <Select
            mode="tags"
            placeholder="Add tags (type to search or create)"
            loading={fetching}
            onSearch={handleTagSearch}
            filterOption={false}
            options={tagOptions.map((t) => ({ label: t.name, value: t.name }))}
            tokenSeparators={[',']}
          />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default CreatePostModal;
