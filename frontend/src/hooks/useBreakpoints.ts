import { useMediaQuery, MOBILE_BREAKPOINT_PX } from '@/hooks/useMediaQuery';

export const TABLET_BREAKPOINT_PX = 1024;
export const SMALL_MOBILE_BREAKPOINT_PX = 576;

/** Viewport buckets for adaptive layout (mobile-first). */
export function useBreakpoints() {
  const { isMobile } = useMediaQuery(MOBILE_BREAKPOINT_PX);
  const { isMobile: isSmallMobile } = useMediaQuery(SMALL_MOBILE_BREAKPOINT_PX);
  const { isMobile: isTabletOrBelow } = useMediaQuery(TABLET_BREAKPOINT_PX);

  return {
    isMobile,
    isSmallMobile,
    isTablet: !isMobile && isTabletOrBelow,
    isDesktop: !isTabletOrBelow,
    /** Alias kept for existing call sites. */
    isTabletOrBelow,
  };
}
