import { UploadCloud, File, X, CheckCircle2 } from 'lucide-react';
import { useState } from 'react';
import axios from 'axios';
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { useToast } from "@/hooks/use-toast";
import { Progress } from "@/components/ui/progress";
import api from './api/axios';
import { endpoints } from './api/endpoints';

interface UploadState {
  progress: number;
  status: 'idle' | 'uploading' | 'success' | 'error';
  fileName: string;
}

function App() {
  const [uploadState, setUploadState] = useState<UploadState>({
    progress: 0,
    status: 'idle',
    fileName: '',
  });
  const { toast } = useToast();

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setUploadState(prev => ({ ...prev, fileName: file.name }));
    }
  };

  const handleUpload = async () => {
    const fileInput = document.querySelector('input[type="file"]') as HTMLInputElement;
    const file = fileInput.files?.[0];

    if (!file) {
      toast({
        title: "Error",
        description: "Please select a file first",
        variant: "destructive",
      });
      return;
    }

    setUploadState(prev => ({ ...prev, status: 'uploading', progress: 0 }));

    try {
      const formData = new FormData();
      formData.append('file', file);

      await api.post(endpoints.upload, formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
        onUploadProgress: (progressEvent) => {
          const progress = progressEvent.total
            ? Math.round((progressEvent.loaded * 100) / progressEvent.total)
            : 0;
          setUploadState((prev) => ({ ...prev, progress }));
        },
      });

      setUploadState(prev => ({ ...prev, status: 'success', progress: 100 }));
      
      toast({
        title: "Success",
        description: "File uploaded successfully",
      });
    } catch (error) {
      setUploadState(prev => ({ ...prev, status: 'error', progress: 0 }));
      toast({
        title: "Error",
        description: axios.isAxiosError(error) 
          ? error.response?.data?.message || "Failed to upload file"
          : "Failed to upload file",
        variant: "destructive",
      });
    }
  };

  const resetUpload = () => {
    setUploadState({
      progress: 0,
      status: 'idle',
      fileName: '',
    });
    const fileInput = document.querySelector('input[type="file"]') as HTMLInputElement;
    if (fileInput) fileInput.value = '';
  };

  return (
    <div className="h-screen w-screen bg-gray-50 flex items-center justify-center p-4">
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <UploadCloud className="h-6 w-6" />
            File Upload
          </CardTitle>
          <CardDescription>
            Upload your files securely to S3 storage
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-4">
            <div className="flex flex-col gap-4">
              <div className="grid w-full items-center gap-2">
                <Input
                  type="file"
                  onChange={handleFileChange}
                  className="cursor-pointer"
                  disabled={uploadState.status === "uploading"}
                />
              </div>

              {uploadState.fileName && (
                <div className="flex items-center gap-2 p-2 bg-gray-50 rounded-md">
                  <File className="h-4 w-4 text-gray-500" />
                  <span className="text-sm text-gray-600 flex-1 truncate">
                    {uploadState.fileName}
                  </span>
                  <button
                    onClick={resetUpload}
                    className="text-gray-500 hover:text-gray-700"
                    disabled={uploadState.status === "uploading"}
                  >
                    <X className="h-4 w-4" />
                  </button>
                </div>
              )}

              {uploadState.status === "uploading" && (
                <div className="space-y-2">
                  <Progress value={uploadState.progress} />
                  <p className="text-sm text-gray-500 text-center">
                    Uploading... {uploadState.progress}%
                  </p>
                </div>
              )}

              {uploadState.status === "success" && (
                <div className="flex items-center gap-2 text-green-600">
                  <CheckCircle2 className="h-4 w-4" />
                  <span className="text-sm">Upload complete!</span>
                </div>
              )}

              <Button
                onClick={handleUpload}
                disabled={
                  !uploadState.fileName || uploadState.status === "uploading"
                }
                className="w-full"
              >
                {uploadState.status === "uploading"
                  ? "Uploading..."
                  : "Upload File"}
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}

export default App;