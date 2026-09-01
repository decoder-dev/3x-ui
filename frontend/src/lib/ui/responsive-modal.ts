import type { CSSProperties } from 'react';
import type { ModalProps } from 'antd';

export type ResponsiveModalOptions = {
  /** Desktop modal width (px or CSS length). */
  desktopWidth?: number | string;
  /** Use edge-to-edge full viewport on phones (default true). */
  fullScreenMobile?: boolean;
};

const defaultBodyStyle = (isMobile: boolean): CSSProperties => ({
  maxHeight: isMobile ? 'calc(100dvh - 108px)' : 'calc(100vh - 160px)',
  overflowY: 'auto',
  overflowX: 'hidden',
  WebkitOverflowScrolling: 'touch',
});

/**
 * Ant Design Modal props for premium adaptive dialogs:
 * desktop — centered card; mobile — full-height sheet with scrollable body.
 */
export function responsiveModalProps(
  isMobile: boolean,
  options: ResponsiveModalOptions = {},
): Pick<ModalProps, 'width' | 'style' | 'className' | 'styles' | 'centered'> {
  const desktopWidth = options.desktopWidth ?? 720;
  const fullScreenMobile = options.fullScreenMobile ?? true;

  if (!isMobile || !fullScreenMobile) {
    return {
      width: isMobile ? 'min(96vw, 560px)' : desktopWidth,
      centered: true,
      styles: { body: defaultBodyStyle(isMobile) },
    };
  }

  return {
    width: '100%',
    centered: false,
    className: 'xui-modal-mobile',
    style: {
      top: 0,
      padding: 0,
      margin: 0,
      maxWidth: '100vw',
    },
    styles: {
      body: {
        ...defaultBodyStyle(true),
        padding: '12px 14px',
        paddingBottom: 'calc(12px + env(safe-area-inset-bottom, 0px))',
      },
      content: {
        borderRadius: 0,
        minHeight: '100dvh',
        maxHeight: '100dvh',
        display: 'flex',
        flexDirection: 'column',
      },
      header: {
        paddingTop: 'calc(12px + env(safe-area-inset-top, 0px))',
        flexShrink: 0,
      },
      footer: {
        flexShrink: 0,
        paddingBottom: 'calc(12px + env(safe-area-inset-bottom, 0px))',
      },
    },
  };
}

/** Horizontal Ant Form: vertical labels on narrow screens. */
export function responsiveFormLayout(isMobile: boolean) {
  if (isMobile) {
    return {
      layout: 'vertical' as const,
      labelCol: undefined,
      wrapperCol: undefined,
    };
  }
  return {
    layout: 'horizontal' as const,
    labelCol: { sm: { span: 8 } },
    wrapperCol: { sm: { span: 14 } },
  };
}
