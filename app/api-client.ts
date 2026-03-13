// Base URL for the Go API.
// In production set NEXT_PUBLIC_API_BASE_URL=https://your-vps-ip:8080
// Leave it unset in development and the Next.js proxy below will forward requests.
const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || '';

export interface VideoInfo {
  title: string;
  thumbnail: string;
  duration?: number;   // seconds
  uploader?: string;
}

export type VideoFormat = 'mp4' | 'webm';
export type VideoQuality = '360' | '480' | '720' | '1080';

export async function getVideoInfo(url: string): Promise<VideoInfo> {
  const response = await fetch(`${API_BASE_URL}/video-info`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ url }),
  });

  if (!response.ok) {
    const { error } = await response.json().catch(() => ({ error: 'Unknown error' }));
    throw new Error(error ?? 'Failed to fetch video info');
  }

  return response.json();
}

/**
 * Returns the download URL. Assign to window.location.href to trigger a
 * browser download, or use as the `src` for a mobile-app fetch stream.
 */
export function getDownloadURL(
  url: string,
  format: VideoFormat,
  quality: VideoQuality,
): string {
  const params = new URLSearchParams({ url, format, quality });
  return `${API_BASE_URL}/download?${params}`;
}

/**
 * Streams the video and returns a Blob — useful for React Native / mobile
 * where you need to save to disk via the file system API.
 */
export async function downloadVideoBlob(
  url: string,
  format: VideoFormat,
  quality: VideoQuality,
  onProgress?: (loaded: number) => void,
): Promise<Blob> {
  const downloadURL = getDownloadURL(url, format, quality);
  const response = await fetch(downloadURL);

  if (!response.ok) {
    throw new Error(`Download failed: ${response.status}`);
  }

  if (!response.body) {
    throw new Error('No response body');
  }

  const reader = response.body.getReader();
  const chunks: Uint8Array<ArrayBuffer>[] = [];
  let loaded = 0;

  while (true) {
    const { done, value } = await reader.read();
    if (done) break;
    chunks.push(value);
    loaded += value.byteLength;
    onProgress?.(loaded);
  }

  const mimeType = format === 'mp4' ? 'video/mp4' : 'video/webm';
  return new Blob(chunks, { type: mimeType });
}