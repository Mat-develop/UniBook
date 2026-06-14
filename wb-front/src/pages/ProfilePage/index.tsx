import React, { useEffect, useRef, useState } from 'react';
import { useParams } from 'react-router-dom';
import {
  Avatar, Button, Card, Col, Form, Input, InputNumber,
  Modal, Row, Spin, Tag, Tooltip, Typography,
} from 'antd';
import {
  DeleteOutlined, EditOutlined, LinkOutlined,
  PlusOutlined, UserOutlined,
} from '@ant-design/icons';
import { toast } from 'react-toastify';
import {
  getProfile, addEducation, updateEducation, deleteEducation,
  addProject, updateProject, deleteProject,
  addCourse, updateCourse, deleteCourse,
  updateUserImage,
  type UserProfile, type Education, type Project, type Course,
} from '../../utils/api';
import { getUserIdFromToken } from '../../utils/auth';
import styles from './profilePage.module.scss';

const { Title, Text } = Typography;

const EMPTY_ED: Education  = { institution: '', degree: '', fieldOfStudy: '', startYear: new Date().getFullYear() };
const EMPTY_PR: Project    = { title: '', description: '' };
const EMPTY_CO: Course     = { title: '', institution: '' };

function resizeImage(file: File, maxPx: number): Promise<string> {
  return new Promise((resolve, reject) => {
    const img = new Image();
    const url = URL.createObjectURL(file);
    img.onload = () => {
      const scale = Math.min(1, maxPx / Math.max(img.width, img.height));
      const canvas = document.createElement('canvas');
      canvas.width  = Math.round(img.width  * scale);
      canvas.height = Math.round(img.height * scale);
      canvas.getContext('2d')!.drawImage(img, 0, 0, canvas.width, canvas.height);
      URL.revokeObjectURL(url);
      resolve(canvas.toDataURL('image/jpeg', 0.85));
    };
    img.onerror = reject;
    img.src = url;
  });
}

export default function ProfilePage() {
  const { userId } = useParams<{ userId: string }>();
  const uid = Number(userId);
  const myId = getUserIdFromToken();
  const isOwner = myId === uid;

  const [profile, setProfile]       = useState<UserProfile | null>(null);
  const [loading, setLoading]       = useState(true);
  const [imgUploading, setImgUploading] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Education modal state
  const [edModal, setEdModal]   = useState(false);
  const [edItem, setEdItem]     = useState<Education>(EMPTY_ED);
  const [edSaving, setEdSaving] = useState(false);

  // Project modal state
  const [prModal, setPrModal]   = useState(false);
  const [prItem, setPrItem]     = useState<Project>(EMPTY_PR);
  const [prSaving, setPrSaving] = useState(false);

  // Course modal state
  const [coModal, setCoModal]   = useState(false);
  const [coItem, setCoItem]     = useState<Course>(EMPTY_CO);
  const [coSaving, setCoSaving] = useState(false);

  const reload = () => getProfile(uid).then(setProfile).catch(() => toast.error('Erro ao carregar perfil'));

  const handleImageChange = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;
    setImgUploading(true);
    try {
      const base64 = await resizeImage(file, 200);
      await updateUserImage(uid, base64);
      setProfile((p) => p ? { ...p, imageUrl: base64 } : p);
    } catch { toast.error('Erro ao atualizar foto'); }
    finally {
      setImgUploading(false);
      if (fileInputRef.current) fileInputRef.current.value = '';
    }
  };

  useEffect(() => {
    getProfile(uid)
      .then(setProfile)
      .catch(() => toast.error('Perfil não encontrado'))
      .finally(() => setLoading(false));
  }, [uid]);

  // ── Education ─────────────────────────────────────────────────────────────
  const openAddEd  = () => { setEdItem(EMPTY_ED); setEdModal(true); };
  const openEditEd = (e: Education) => { setEdItem(e); setEdModal(true); };
  const saveEd = async () => {
    setEdSaving(true);
    try {
      if (edItem.id) await updateEducation(edItem.id, edItem);
      else await addEducation(edItem);
      setEdModal(false);
      reload();
    } catch { toast.error('Erro ao salvar formação'); }
    finally { setEdSaving(false); }
  };
  const removeEd = async (id: number) => {
    try { await deleteEducation(id); reload(); }
    catch { toast.error('Erro ao remover formação'); }
  };

  // ── Projects ──────────────────────────────────────────────────────────────
  const openAddPr  = () => { setPrItem(EMPTY_PR); setPrModal(true); };
  const openEditPr = (p: Project) => { setPrItem(p); setPrModal(true); };
  const savePr = async () => {
    setPrSaving(true);
    try {
      if (prItem.id) await updateProject(prItem.id, prItem);
      else await addProject(prItem);
      setPrModal(false);
      reload();
    } catch { toast.error('Erro ao salvar projeto'); }
    finally { setPrSaving(false); }
  };
  const removePr = async (id: number) => {
    try { await deleteProject(id); reload(); }
    catch { toast.error('Erro ao remover projeto'); }
  };

  // ── Courses ───────────────────────────────────────────────────────────────
  const openAddCo  = () => { setCoItem(EMPTY_CO); setCoModal(true); };
  const openEditCo = (c: Course) => { setCoItem(c); setCoModal(true); };
  const saveCo = async () => {
    setCoSaving(true);
    try {
      if (coItem.id) await updateCourse(coItem.id, coItem);
      else await addCourse(coItem);
      setCoModal(false);
      reload();
    } catch { toast.error('Erro ao salvar curso'); }
    finally { setCoSaving(false); }
  };
  const removeCo = async (id: number) => {
    try { await deleteCourse(id); reload(); }
    catch { toast.error('Erro ao remover curso'); }
  };

  if (loading) return <div className={styles.center}><Spin size="large" /></div>;
  if (!profile) return <div className={styles.center}><Text type="danger">Perfil não encontrado.</Text></div>;

  return (
    <div className={styles.page}>

      {/* ── Header ─────────────────────────────────────────────────────── */}
      <Card className={styles.headerCard}>
        <div className={styles.headerInner}>
          <Tooltip title={isOwner ? 'Alterar foto' : undefined}>
            <div
              className={`${styles.avatarWrapper} ${isOwner ? styles.avatarOwner : ''}`}
              onClick={() => isOwner && fileInputRef.current?.click()}
            >
              <Avatar size={80} src={profile.imageUrl || undefined} icon={<UserOutlined />} />
              {isOwner && <div className={styles.avatarOverlay}>{imgUploading ? <Spin size="small" /> : '📷'}</div>}
            </div>
          </Tooltip>
          <input
            ref={fileInputRef}
            type="file"
            accept="image/*"
            style={{ display: 'none' }}
            onChange={handleImageChange}
          />
          <div>
            <Title level={3} className={styles.name}>{profile.name}</Title>
            <Text type="secondary">@{profile.nick}</Text>
          </div>
        </div>
      </Card>

      <Row gutter={20} className={styles.body}>

        {/* ── Educação (coluna esquerda) ──────────────────────────────── */}
        <Col xs={24} md={8}>
          <Card
            title="Formação Acadêmica"
            className={styles.card}
            extra={isOwner && <Button size="small" icon={<PlusOutlined />} onClick={openAddEd} />}
          >
            {profile.education.length === 0
              ? <Text type="secondary">Nenhuma formação cadastrada.</Text>
              : profile.education.map((e) => (
                  <div key={e.id} className={styles.item}>
                    <div className={styles.itemHeader}>
                      <Text strong>{e.degree}</Text>
                      {isOwner && (
                        <span className={styles.actions}>
                          <Tooltip title="Editar"><EditOutlined onClick={() => openEditEd(e)} /></Tooltip>
                          <Tooltip title="Remover"><DeleteOutlined onClick={() => removeEd(e.id!)} /></Tooltip>
                        </span>
                      )}
                    </div>
                    <Text>{e.fieldOfStudy}</Text>
                    <br />
                    <Text type="secondary">{e.institution}</Text>
                    <br />
                    <Text type="secondary" className={styles.period}>
                      {e.startYear} — {e.endYear ?? 'presente'}
                    </Text>
                    {e.description && <p className={styles.desc}>{e.description}</p>}
                  </div>
                ))
            }
          </Card>
        </Col>

        {/* ── Projetos + Cursos (coluna direita) ─────────────────────── */}
        <Col xs={24} md={16}>

          {/* Projetos Acadêmicos */}
          <Card
            title="Projetos Acadêmicos"
            className={styles.card}
            extra={isOwner && <Button size="small" icon={<PlusOutlined />} onClick={openAddPr} />}
          >
            {profile.projects.length === 0
              ? <Text type="secondary">Nenhum projeto cadastrado.</Text>
              : profile.projects.map((p) => (
                  <div key={p.id} className={styles.item}>
                    <div className={styles.itemHeader}>
                      <Text strong>{p.title}</Text>
                      <span className={styles.meta}>
                        {p.year && <Tag>{p.year}</Tag>}
                        {isOwner && (
                          <span className={styles.actions}>
                            <Tooltip title="Editar"><EditOutlined onClick={() => openEditPr(p)} /></Tooltip>
                            <Tooltip title="Remover"><DeleteOutlined onClick={() => removePr(p.id!)} /></Tooltip>
                          </span>
                        )}
                      </span>
                    </div>
                    <p className={styles.desc}>{p.description}</p>
                    {p.url && (
                      <a href={p.url} target="_blank" rel="noopener noreferrer" className={styles.link}>
                        <LinkOutlined /> {p.url}
                      </a>
                    )}
                  </div>
                ))
            }
          </Card>

          {/* Cursos */}
          <Card
            title="Cursos e Certificações"
            className={`${styles.card} ${styles.cardTop}`}
            extra={isOwner && <Button size="small" icon={<PlusOutlined />} onClick={openAddCo} />}
          >
            {profile.courses.length === 0
              ? <Text type="secondary">Nenhum curso cadastrado.</Text>
              : <div className={styles.courseGrid}>
                  {profile.courses.map((c) => (
                    <div key={c.id} className={styles.courseItem}>
                      <div className={styles.itemHeader}>
                        <Text strong>{c.title}</Text>
                        {isOwner && (
                          <span className={styles.actions}>
                            <Tooltip title="Editar"><EditOutlined onClick={() => openEditCo(c)} /></Tooltip>
                            <Tooltip title="Remover"><DeleteOutlined onClick={() => removeCo(c.id!)} /></Tooltip>
                          </span>
                        )}
                      </div>
                      <Text type="secondary">{c.institution}{c.year ? ` · ${c.year}` : ''}</Text>
                      {c.url && (
                        <a href={c.url} target="_blank" rel="noopener noreferrer" className={styles.link}>
                          <LinkOutlined /> Certificado
                        </a>
                      )}
                    </div>
                  ))}
                </div>
            }
          </Card>

        </Col>
      </Row>

      {/* ── Modal: Formação ────────────────────────────────────────────────── */}
      <Modal
        title={edItem.id ? 'Editar Formação' : 'Adicionar Formação'}
        open={edModal}
        onOk={saveEd}
        onCancel={() => setEdModal(false)}
        confirmLoading={edSaving}
        okText="Salvar"
      >
        <Form layout="vertical">
          <Form.Item label="Instituição" required>
            <Input value={edItem.institution} onChange={(e) => setEdItem({ ...edItem, institution: e.target.value })} />
          </Form.Item>
          <Form.Item label="Grau / Título" required>
            <Input placeholder="ex: Bacharel, Mestre, Técnico" value={edItem.degree} onChange={(e) => setEdItem({ ...edItem, degree: e.target.value })} />
          </Form.Item>
          <Form.Item label="Área de estudo" required>
            <Input value={edItem.fieldOfStudy} onChange={(e) => setEdItem({ ...edItem, fieldOfStudy: e.target.value })} />
          </Form.Item>
          <Row gutter={12}>
            <Col span={12}>
              <Form.Item label="Ano de início" required>
                <InputNumber style={{ width: '100%' }} value={edItem.startYear} onChange={(v) => setEdItem({ ...edItem, startYear: v ?? 0 })} />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item label="Ano de conclusão">
                <InputNumber style={{ width: '100%' }} placeholder="Em andamento" value={edItem.endYear ?? undefined} onChange={(v) => setEdItem({ ...edItem, endYear: v ?? null })} />
              </Form.Item>
            </Col>
          </Row>
          <Form.Item label="Descrição">
            <Input.TextArea rows={3} value={edItem.description ?? ''} onChange={(e) => setEdItem({ ...edItem, description: e.target.value || null })} />
          </Form.Item>
        </Form>
      </Modal>

      {/* ── Modal: Projeto ─────────────────────────────────────────────────── */}
      <Modal
        title={prItem.id ? 'Editar Projeto' : 'Adicionar Projeto'}
        open={prModal}
        onOk={savePr}
        onCancel={() => setPrModal(false)}
        confirmLoading={prSaving}
        okText="Salvar"
      >
        <Form layout="vertical">
          <Form.Item label="Título" required>
            <Input value={prItem.title} onChange={(e) => setPrItem({ ...prItem, title: e.target.value })} />
          </Form.Item>
          <Form.Item label="Descrição" required>
            <Input.TextArea rows={4} value={prItem.description} onChange={(e) => setPrItem({ ...prItem, description: e.target.value })} />
          </Form.Item>
          <Row gutter={12}>
            <Col span={12}>
              <Form.Item label="Ano">
                <InputNumber style={{ width: '100%' }} value={prItem.year ?? undefined} onChange={(v) => setPrItem({ ...prItem, year: v ?? null })} />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item label="Link">
                <Input prefix={<LinkOutlined />} placeholder="https://..." value={prItem.url ?? ''} onChange={(e) => setPrItem({ ...prItem, url: e.target.value || null })} />
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </Modal>

      {/* ── Modal: Curso ───────────────────────────────────────────────────── */}
      <Modal
        title={coItem.id ? 'Editar Curso' : 'Adicionar Curso'}
        open={coModal}
        onOk={saveCo}
        onCancel={() => setCoModal(false)}
        confirmLoading={coSaving}
        okText="Salvar"
      >
        <Form layout="vertical">
          <Form.Item label="Nome do curso" required>
            <Input value={coItem.title} onChange={(e) => setCoItem({ ...coItem, title: e.target.value })} />
          </Form.Item>
          <Form.Item label="Instituição" required>
            <Input value={coItem.institution} onChange={(e) => setCoItem({ ...coItem, institution: e.target.value })} />
          </Form.Item>
          <Row gutter={12}>
            <Col span={12}>
              <Form.Item label="Ano de conclusão">
                <InputNumber style={{ width: '100%' }} value={coItem.year ?? undefined} onChange={(v) => setCoItem({ ...coItem, year: v ?? null })} />
              </Form.Item>
            </Col>
            <Col span={12}>
              <Form.Item label="Link do certificado">
                <Input prefix={<LinkOutlined />} placeholder="https://..." value={coItem.url ?? ''} onChange={(e) => setCoItem({ ...coItem, url: e.target.value || null })} />
              </Form.Item>
            </Col>
          </Row>
        </Form>
      </Modal>

    </div>
  );
}
