# Github-api-service

<br>
This GO service utilises Github API routes 

<h2>Routes</h2> 
<h3><U>Repository Exists</U></h3> 
Endpoint: /api/repository/exists<br>
Required: 
<ul><li>repoName: Github-api-service</li> 
	<li>user: kalem</li></ul>
Api Route: https://api.github.com/repos/EAS-$user/$repoName<br>
<br>


<h3><U>Repository Contains File</U></h3> 
Endpoint: /api/repository/contains<br>
Required: 
<ul><li>repoName: Github-api-service</li> 
	<li>user: kalem</li>
	<li>fileName: README.md</li>
    </li> </ul>
Api Route: https://api.github.com/repos/EAS-$user/$repoName/contents<br>
<br>


<h3><U>File In Repository Contains</U></h3> 
Endpoint: /api/repository/file/contains<br>
Required: 
<ul><li>repoName: Github-api-service</li> 
	<li>user: kalem</li>
	<li>fileName: README.md</li>
    <li>check: # Github-api-service</li> </ul>
Api Route: https://api.github.com/repos/EAS-$user/$repoName/contents/$fileName<br>
<br>
<h2>TODO:</h2>
<ul><li>Clean code</li> 
	<li>Utilise models folder as modules</li>
</ul># Vmware-api-service
