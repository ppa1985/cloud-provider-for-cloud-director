/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul> 
 *
 * API version: 37.2
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// The VCD SSL settings
type SslSettings struct {
	// SSL protocols
	EnabledSslProtocols []string `json:"enabledSslProtocols,omitempty"`
	// SSL ciphers
	EnabledSslCiphers []string `json:"enabledSslCiphers,omitempty"`
	// Size of RSA keys generated by VCD. In some circumstances, will also be the minimum key size that will be enforced by certain features and will be so noted in respective documentation.
	KeySize int32 `json:"keySize,omitempty"`
	// Size of EC keys generated by VCD. In some circumstances, will also be the minimum key size that will be enforced by certain features and will be so noted in respective documentation.
	EcKeySize int32 `json:"ecKeySize,omitempty"`
	// Number of days any VCD generated certificate will be valid for
	CertificateValidityDays int32 `json:"certificateValidityDays,omitempty"`
	// Algorithm VCD will use to sign any self-generated certificates
	CertificateSignatureAlgorithm string `json:"certificateSignatureAlgorithm,omitempty"`
	// The desired FIPS mode of this server group. <ul>   <li>     <em>ON</em>: Indicates FIPS mode is desired for this server group.   </li>   <li>     <em>OFF</em>: Indicates FIPS mode is not desired for this server group.   </li> </ul> 
	FipsMode string `json:"fipsMode,omitempty"`
	// The current SSL settings status for this server group. <ul>   <li>     <em>CURRENT</em>: Indicates that all of the SSL settings for this server group are up to date.   </li>   <li>     <em>UPDATING</em>: Indicates that there is at least one cell in the server group which has not yet     applied the SSL settings.   </li>   <li>     <em>AWAITING_RESTART</em>: Indicates that the SSL settings for this server group have been applied on each cell,     and that each cell needs to be rebooted for the settings to take place. When each cell has been rebooted this field     will change to CURRENT.   </li> </ul> 
	Status string `json:"status,omitempty"`
}
