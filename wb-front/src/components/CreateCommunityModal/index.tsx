import React, { useState } from 'react';
import { Modal, Form, Input, Button } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { toast } from 'react-toastify';
import { createCommunity } from '../../utils/api';

interface Props {
  onCreated?: () => void;
}

const CreateCommunityModal: React.FC<Props> = ({ onCreated }) => {
  const [open, setOpen] = useState(false);
  const [saving, setSaving] = useState(false);
  const [form] = Form.useForm();

  const handleOk = async () => {
    try {
      const values = await form.validateFields();
      setSaving(true);
      await createCommunity(values);
      toast.success('Comunidade criada com sucesso!');
      form.resetFields();
      setOpen(false);
      onCreated?.();
    } catch (err: unknown) {
      const msg = (err as { response?: { data?: { erro?: string } } })?.response?.data?.erro;
      if (msg) toast.error(msg);
    } finally {
      setSaving(false);
    }
  };

  return (
    <>
      <Button type="primary" icon={<PlusOutlined />} onClick={() => setOpen(true)}>
        Criar Comunidade
      </Button>

      <Modal
        title="Criar nova comunidade"
        open={open}
        onOk={handleOk}
        onCancel={() => { setOpen(false); form.resetFields(); }}
        confirmLoading={saving}
        okText="Criar"
        cancelText="Cancelar"
      >
        <Form form={form} layout="vertical">
          <Form.Item name="name" label="Nome" rules={[{ required: true, message: 'Informe o nome' }]}>
            <Input placeholder="ex: Engenharia de Software" />
          </Form.Item>
          <Form.Item name="description" label="Descrição" rules={[{ required: true, message: 'Informe a descrição' }]}>
            <Input.TextArea rows={3} placeholder="Do que se trata essa comunidade?" />
          </Form.Item>
          <Form.Item name="imageUrl" label="URL da imagem (opcional)">
            <Input placeholder="https://..." />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default CreateCommunityModal;
