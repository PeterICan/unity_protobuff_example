using UnityEditor;
using UnityEngine;
using System.IO;
using UnityEditor.Build.Reporting; 
using UnityEditor.Build.Reporting; 

public class BuildScript
{
    // 這個方法將透過命令列呼叫
    public static void PerformBuild()
    {
        // 1. 定義輸出路徑和檔案名稱
        string outputDir = "Builds";
        string productName = PlayerSettings.productName;
        string buildPath = Path.Combine(Application.dataPath, "..", outputDir, productName + "_Windows.exe");

        // 確保輸出目錄存在
        if (!Directory.Exists(Path.GetDirectoryName(buildPath)))
        {
            Directory.CreateDirectory(Path.GetDirectoryName(buildPath));
        }

        // 2. 設定建置選項 (BuildPlayerOptions)
        BuildPlayerOptions buildOptions = new BuildPlayerOptions();
        
        // 設定要包含在建置中的場景 (這裡使用 EditorBuildSettings 中的場景清單)
        buildOptions.scenes = GetScenePaths();
        
        // 設定輸出路徑
        buildOptions.locationPathName = buildPath;
        
        // 設定目標平台 (例如 Windows 64-bit)
        buildOptions.target = BuildTarget.StandaloneWindows64; 
        
        // 設定建置選項 (例如：BuildOptions.None 或 BuildOptions.Development)
        buildOptions.options = BuildOptions.None; 

        // 3. 執行建置
        BuildReport report = BuildPipeline.BuildPlayer(buildOptions);

        // 4. 處理建置結果
        if (report.summary.result == BuildResult.Succeeded)
        {
            Debug.Log($"[CLI Build SUCCESS]: Path: {buildPath}");
        }
        else if (report.summary.result == BuildResult.Failed)
        {
            Debug.LogError("[CLI Build FAILED]");
        }
        else
        {
            Debug.LogWarning($"[CLI Build Finished with {report.summary.result}]");
        }
    }

    // 輔助方法：取得 EditorBuildSettings 中所有啟用的場景路徑
    private static string[] GetScenePaths()
    {
        string[] scenes = new string[EditorBuildSettings.scenes.Length];
        for (int i = 0; i < EditorBuildSettings.scenes.Length; i++)
        {
            scenes[i] = EditorBuildSettings.scenes[i].path;
        }
        return scenes;
    }
}