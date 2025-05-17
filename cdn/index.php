<?php
header_remove("X-Powered-By");

abstract class FolderScanner
{
    public string $path;
    private array|false $directories;
    private array $folders = [];

    public function __construct(string $path)
    {

        $this->path = realpath(sprintf("./%s/", $path));

        if (!$this->path) {
            mkdir(sprintf("./%s/", $path));
            $this->path = realpath(sprintf("./%s/", $path));
        };

        $this->directories = scandir($this->path, SCANDIR_SORT_ASCENDING);
        $this->start();
    }

    private function start(): void
    {
        if ($this->directories) {
            foreach ($this->directories as $dir) {

                $pattern = '/\W/';
                if (preg_match($pattern, $dir)) continue;

                if (is_dir($this->path . '/' . $dir)) {
                    $this->folders[] = strtolower($dir);
                }
            }
        }
    }

    protected function getFolders(): array
    {
        return $this->folders;
    }
}

class ContentDeliveryNetwork extends FolderScanner
{
    private string $uri;

    public function __construct(string $mainFolder, string $request_uri)
    {
        parent::__construct($mainFolder);

        $this->uri = $request_uri;

        $this->handleRequest();
    }

    private function handleRequest(): void
    {
        $uri = explode("/", $this->uri);

        if (count($uri) < 3) {
            $this->not_found();
            return;
        }

        $path = $uri[1] ?? '';
        $image = $uri[2] ?? '';

        if (in_array(strtolower($path), $this->getFolders())) {
            $this->load_image(sprintf("%s/%s/%s", $this->path, strtolower($path), $image));
            return;
        }

        $this->not_found();
    }

    private function load_image(string $fileName): void
    {
        if (!file_exists($fileName)) {
            http_response_code(404);
            header("Content-type: application/json");
            echo json_encode(array("Cod" => 404, "Msg" => "image.not.found"));
            return;
        }

        $this->header($fileName);
        readfile($fileName);
    }

    private function header(string $fileName): void
    {
        header("Content-Type:" . mime_content_type($fileName));
        header("Cache-Control: public, max-age=31536000, must-revalidate");
        header("Content-Length: " . filesize($fileName));
        header("X-Robots-Tag: noindex, nofollow, noarchive, nocache, noimageindex, noodp");
    }

    private function not_found(): void
    {
        http_response_code(404);
        header("Content-type: application/json");
        echo json_encode(array("Cod" => 404, "Msg" => "not.found", "Routes" => join(", ", $this->getFolders())));
    }
}

new ContentDeliveryNetwork("assets", $_SERVER["REQUEST_URI"]);
