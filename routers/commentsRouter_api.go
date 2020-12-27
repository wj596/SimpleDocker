package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/api/container",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "RemoveContainer",
            Router: "/api/container/:containerId/delete",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "GetContainerInfo",
            Router: "/api/container/:containerId/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "GetContainerLog",
            Router: "/api/container/:containerId/log",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "GetContainerAllLog",
            Router: "/api/container/:containerId/log/all",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "RestartContainer",
            Router: "/api/container/:containerId/restart",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "StartContainer",
            Router: "/api/container/:containerId/start",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "StopContainer",
            Router: "/api/container/:containerId/stop",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("containerId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ContainerController"],
        beego.ControllerComments{
            Method: "CreateNewContainer",
            Router: "/api/container/run",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:DockerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:DockerController"],
        beego.ControllerComments{
            Method: "DockerInfo",
            Router: "/api/docker/info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:DockerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:DockerController"],
        beego.ControllerComments{
            Method: "Ping",
            Router: "/api/docker/ping",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:DockerController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:DockerController"],
        beego.ControllerComments{
            Method: "GetVersion",
            Router: "/api/docker/version",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ImageController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ImageController"],
        beego.ControllerComments{
            Method: "GetImageList",
            Router: "/api/image",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ImageController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ImageController"],
        beego.ControllerComments{
            Method: "GetImageInfo",
            Router: "/api/image/:imageId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("imageId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ImageController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ImageController"],
        beego.ControllerComments{
            Method: "DeleteImage",
            Router: "/api/image/:imageId/remove/:forge",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("imageId", param.InPath),
				param.New("forge", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ImageController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ImageController"],
        beego.ControllerComments{
            Method: "SaveImage",
            Router: "/api/image/:imageId/save",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("imageId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:ImageController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:ImageController"],
        beego.ControllerComments{
            Method: "TagImage",
            Router: "/api/image/tag",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"],
        beego.ControllerComments{
            Method: "GetVolumeList",
            Router: "/api/volume",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"],
        beego.ControllerComments{
            Method: "NewVolume",
            Router: "/api/volume/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(
				param.New("name", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"],
        beego.ControllerComments{
            Method: "GetVolumeInfo",
            Router: "/api/volume/:volumeId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(
				param.New("volumeId", param.InPath),
			),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"] = append(beego.GlobalControllerRouter["SimpleDocker/api:VolumeController"],
        beego.ControllerComments{
            Method: "RemoveVolume",
            Router: "/api/volume/:volumeId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(
				param.New("volumeId", param.InPath),
			),
            Filters: nil,
            Params: nil})

}
